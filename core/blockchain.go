package core

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain_%s.db"
const blocksBucket = "blocks"

type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

// // AddBlock adds a block to the blockchain
// func (bc *BlockChain) AddBlock(block *Block) {
// 	prevBlock := bc.blocks[len(bc.blocks)-1]
// 	// newBlock := NewBlock(data, prevBlock.Hash)
// 	// bc.blocks = append(bc.blocks, newBlock)

// }

// // NewGenesisBlock creates and returns genesis Block
// func NewGenesisBlock() *Block {
// 	return NewBlock("Genesis Block", []byte{})
// }

// // NewBlockChain creates a new BlockChain with genesis Block
// func NewBlockChain() *BlockChain {
// 	return &BlockChain{[]*Block{NewGenesisBlock()}}
// }

// FindTransaction finds a transaction by its ID
func (bc *BlockChain) FindTransaction(ID []byte) (Transaction, error) {
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, tx := range block.Body.Transactions {
			if bytes.Compare(tx.ID, ID) == 0 {
				return *tx, nil
			}
		}

		if len(block.Header.PrevBlock) == 0 {
			break
		}
	}

	return Transaction{}, errors.New("Transaction is not found")
}

// FindUTXO finds all unspent transaction outputs and returns transactions with spent outputs removed
func (bc *BlockChain) FindUTXO() map[string]TXOutputs {
	UTXO := make(map[string]TXOutputs)
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, tx := range block.Body.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {
				// Was the output spent?
				if spentTXOs[txID] != nil {
					for _, spentOutIdx := range spentTXOs[txID] {
						if spentOutIdx == outIdx {
							continue Outputs
						}
					}
				}

				outs := UTXO[txID]
				outs.Outputs = append(outs.Outputs, out)
				UTXO[txID] = outs
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					inTxID := hex.EncodeToString(in.Txid)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
				}
			}
		}

		if len(block.Header.PrevBlock) == 0 {
			break
		}
	}

	return UTXO
}

// Iterator returns a BlockchainIterat
func (bc *BlockChain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip, bc.db}

	return bci
}

// SignTransaction signs inputs of a Transaction
func (bc *BlockChain) SignTransaction(tx *Transaction, privKey *ecdsa.PrivateKey) {
	prevTXs := make(map[string]Transaction)

	for _, vin := range tx.Vin {
		prevTX, err := bc.FindTransaction(vin.Txid)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}

	tx.Sign(privKey, prevTXs)
}
