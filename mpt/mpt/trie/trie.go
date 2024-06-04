package trie

import (
	"bytes"
	"errors"

	"math/big"
	"mpt/crypto/sha3"
	"mpt/kvstore"
	"mpt/types"
	"mpt/utils/hash"
	"mpt/utils/hexutil"
	"mpt/utils/rlp"
	"sort"
	"strings"
)

var EmptyHash = hash.BigToHash(big.NewInt(0))

type (
	ITrie interface {
		Store(key []byte, account types.Account) error
		Root() hash.Hash
		Load(key []byte) (types.Account, error)
	}
	Child struct {
		Path string
		Hash hash.Hash
	}
	Children []Child

	TrieNode struct {
		Path      string    //节点路径
		Leaf      bool      //是否是叶子节点
		ValueHash hash.Hash //当前节点的哈希值
		Children  Children  //孩子类型的切片
	}
	State struct {
		db   kvstore.KVDatabase
		root *TrieNode
	}
)

func NewChild(path string, hash hash.Hash) Child {
	return Child{
		Path: path,
		Hash: hash,
	}
}

func (children Children) Len() int {
	return len(children)
}

func (children Children) Less(i, j int) bool {
	return strings.Compare(children[i].Path, children[j].Path) < 0

}
func (children Children) Swap(i, j int) {
	children[i], children[j] = children[j], children[i]
}
func (node *TrieNode) Sort() {
	sort.Sort(node.Children)
}

func NewTrienode() *TrieNode {
	return &TrieNode{
		Path: "",
	}
}

func TrieNodeFromBytes(data []byte) (*TrieNode, error) {
	var node TrieNode
	err := rlp.DecodeBytes(data, &node)
	return &node, err

}

func NewState(db kvstore.KVDatabase, root hash.Hash) *State {
	if bytes.Equal(EmptyHash[:], root[:]) {
		return &State{
			db:   db,
			root: NewTrienode(),
		}
	} else {
		value, err := db.Get(root[:])
		if err != nil {
			panic(err)
		}
		node, err := TrieNodeFromBytes(value)
		if err != nil {
			panic(err)
		}
		return &State{
			db:   db,
			root: node,
		}
	}

}

func (node *TrieNode) Bytes() []byte {
	data, _ := rlp.EncodeToBytes(node)
	return data
}
func (node *TrieNode) Hash() hash.Hash {
	data := node.Bytes()
	return sha3.Keccak256(data)
}
func (state *State) Root() hash.Hash {
	return state.root.Hash()
}

// func (state *State) pri() {
// 	fmt.Println(state.root.Leaf)
// 	fmt.Println(state.root.Path)
// 	fmt.Println(state.root.ValueHash)
// }

func prefixLength(s1, s2 string) int {
	length := len(s1)
	if length > len(s2) {
		return len(s2)
	}
	for i := 0; i < length; i++ {
		if s1[i] != s2[i] {
			return i
		}
	}
	return length
}

func (state *State) FindAncestors(path string) ([]string, []hash.Hash) { //查找
	current := state.root
	paths, hashs := make([]string, 0), make([]hash.Hash, 0)
	paths = append(paths, "")
	hashs = append(hashs, state.Root())
	prefix := ""
	for {
		flag := false
		for _, child := range current.Children {
			tmp := prefix + child.Path
			length := prefixLength(path, tmp)
			if length == len(tmp) {
				prefix = prefix + child.Path
				paths = append(paths, child.Path)
				hashs = append(hashs, child.Hash)
				flag = true
				data, _ := state.db.Get(child.Hash[:])
				current, _ = TrieNodeFromBytes(data)
				break
			} else if length > len(prefix) {
				l := length - len(prefix)
				str := child.Path[:l]
				paths = append(paths, str)
				hashs = append(hashs, child.Hash)
				return paths, hashs
			}

		}
		if !flag {
			break
		}
	}
	return paths, hashs
}

func (state *State) LoadTeieNodeFromHash(hash hash.Hash) (*TrieNode, error) {
	data, err := state.db.Get(hash[:])
	if err != nil {
		return nil, err
	}
	return TrieNodeFromBytes(data)
}
func (state *State) Save(node TrieNode) {
	h := node.Hash()
	state.db.Put(h[:], node.Bytes())
}

func (state *State) Load(key []byte) (types.Account, error) {
	path := hexutil.Encode(key)
	path = path[2:]
	paths, hashs := state.FindAncestors(path)
	matched := strings.Join(paths, "")
	var account types.Account
	if strings.EqualFold(path, matched) {
		lastHash := hashs[len(hashs)-1]
		leafNode, err := state.LoadTeieNodeFromHash(lastHash)
		if err != nil {
			return account, err
		}
		if !leafNode.Leaf {
			return account, errors.New("not found")

		}
		data, err := state.db.Get(leafNode.ValueHash[:])
		_ = rlp.DecodeBytes(data, &account)
		return account, err
	} else {
		return account, errors.New("哦，没找到，玩球，代码出错了")
	}

}
func (state *State) Update(node *TrieNode, hash []hash.Hash) {
	childHash := node.Hash()
	childPath := node.Path
	depth := len(hash)
	if depth == 1 {
		state.root = node
	}
	for i := depth - 2; i >= 0; i-- {
		current, _ := state.LoadTeieNodeFromHash(hash[i]) //
		for key := range current.Children {
			if strings.Contains(current.Children[key].Path, childPath) {
				current.Children[key].Hash = childHash
				current.Children[key].Path = childPath
				state.Save(*current)
				childHash = current.Hash()
				childPath = current.Path
				break
			}
		}
		if i == 0 {
			state.root = current
		}
	}

}

func (state *State) Store(key []byte, account types.Account) error {
	value := account.Bytes()
	valueHash := sha3.Keccak256(value)
	state.db.Put(valueHash[:], value)
	path := hexutil.Encode(key)
	path = path[2:]
	paths, hashs := state.FindAncestors(path)
	prefix := strings.Join(paths, "")
	depth := len(hashs)
	node, _ := state.LoadTeieNodeFromHash(hashs[depth-1])
	if strings.EqualFold(prefix, path) {
		node.ValueHash = valueHash
		state.Save(*node)
		state.Update(node, hashs)

	} else {
		if strings.EqualFold(node.Path, paths[depth-1]) {
			prefix := strings.Join(paths, "")
			leafPath := path[len(prefix):] //想要插入的节点
			leafNode := NewTrienode()
			leafNode.Leaf = true
			leafNode.Path = leafPath
			leafNode.ValueHash = valueHash
			state.Save(*leafNode)
			leafHash := leafNode.Hash()
			node.Children = append(node.Children, NewChild(leafPath, leafHash))
			node.Sort()
			state.Save(*node)
			state.Update(node, hashs)
		} else {
			//第一步
			lastMatched := paths[len(paths)-1]
			node.Path = node.Path[len(lastMatched):]
			state.Save(*node)
			prefix := strings.Join(paths, "")
			leafPath := path[len(prefix):]
			//第二步
			leafNode := NewTrienode()
			leafNode.Leaf = true
			leafNode.Path = leafPath
			leafNode.ValueHash = valueHash
			state.Save(*leafNode)
			//孩子的父亲
			newNode := NewTrienode()
			newNode.Path = lastMatched
			newNode.Children = make(Children, 0)
			newNode.Children = append(newNode.Children, NewChild(node.Path, node.Hash()), NewChild(leafNode.Path, leafNode.Hash()))
			newNode.Sort()
			state.Save(*newNode)
			state.Update(newNode, hashs)
		}
	}
	return nil
}
