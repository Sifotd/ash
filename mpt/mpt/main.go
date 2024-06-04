package main

import (
	"fmt"
	"mpt/kvstore"
	"mpt/trie"
	"mpt/types"
	"mpt/utils/hexutil"
)

func main() {
	db := kvstore.NewLevelDB("./leveldb")
	state := trie.NewState(db, trie.EmptyHash)
	state.Save(*trie.NewTrienode())
	address, _ := hexutil.Decode("0xB26feabc")
	account1 := types.Account{
		Amount: 999,
		Nonce:  120,
	}
	state.Store(address, account1) //0x68656c6c6f,
	account2, _ := state.Load(address)
	fmt.Println(account2)
	address1, _ := hexutil.Decode("0xB26feabd")
	account3 := types.Account{
		Amount: 998,
		Nonce:  123,
	}
	state.Store(address1, account3)

	account4, _ := state.Load(address1)
	fmt.Println(account4)
	//bulid tx
	//verify
	//change
	//update state
	//load account
}
