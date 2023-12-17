package core

import (
	"github.com/ethereum/go-ethereum/common"
)

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte // Tx Hash
	Vout      int    // Previous Tx Output Index
	Signature []byte
	Address   common.Address
}

// // UsesKey checks whether the address initiated the transaction
// func (in *TXInput) UsesKey(pubKeyByte []byte) bool {
// 	pk, _ := crypto.UnmarshalPubkey(pubKeyByte)
// 	lockingHash := crypto.PubkeyToAddress(*pk)

// 	return bytes.Compare(lockingHash[:], in.Address[:]) == 0
// }
