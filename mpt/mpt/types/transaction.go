package types

import (
	"fmt"
	"math/big"
	"mpt/crypto/secp256k1"
	"mpt/crypto/sha3"
	"mpt/utils/rlp"
)

type Transaction struct {
	txdata
	signature
}

type txdata struct {
	To       Address
	Nonce    uint64
	Value    uint64
	Gas      uint64
	GasPrice uint64
	Input    []byte
}

type signature struct {
	R, S *big.Int
	v    uint8
}

func (tx Transaction) From() Address {
	txdata := tx.txdata
	toSign, err := rlp.EncodeToBytes(txdata)
	fmt.Println(toSign, err)
	msg := sha3.Keccak256(toSign)
	sig := make([]byte, 65)
	pubKey, err := secp256k1.RecoverPubkey(msg[:], sig)
	if err != nil {
		panic(err)
	}
	return PubKeyToAddress(pubKey)
}
