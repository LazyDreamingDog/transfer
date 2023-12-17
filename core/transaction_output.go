package core

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// TXOutput represents a transaction output
type TXOutput struct {
	Value   int
	Address common.Address
}

// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address common.Address) *TXOutput {
	txo := &TXOutput{value, address}
	return txo
}

// Lock signs the output
func (out *TXOutput) Lock(address common.Address) {
	out.Address = address
}

// IsLockedWithKey checks if the output can be used by the owner of the pubkey
func (out *TXOutput) IsLockedWithKey(pubKeyByte []byte) bool {
	pk, _ := crypto.UnmarshalPubkey(pubKeyByte)
	addr := crypto.PubkeyToAddress(*pk)
	return bytes.Compare(out.Address[:], addr[:]) == 0
}

// TXOutputs collects TXOutput
type TXOutputs struct {
	Outputs []TXOutput
}

// Serialize serializes TXOutputs
func (outs TXOutputs) Serialize() []byte {
	var buff bytes.Buffer

	enc := gob.NewEncoder(&buff)
	err := enc.Encode(outs)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

// DeserializeOutputs deserializes TXOutputs
func DeserializeOutputs(data []byte) TXOutputs {
	var outputs TXOutputs

	dec := gob.NewDecoder(bytes.NewReader(data))
	err := dec.Decode(&outputs)
	if err != nil {
		log.Panic(err)
	}

	return outputs
}
