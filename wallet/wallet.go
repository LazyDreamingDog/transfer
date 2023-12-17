package wallet

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// const version = byte(0x00)
// const addressChecksumLen = 4

// Wallet stores private and public keys
type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	PublicKey  ecdsa.PublicKey
	// Address common.Address
}

// NewWallet creates and returns a Wallet
func NewWallet() *Wallet {
	private, _ := crypto.GenerateKey()
	// addr := crypto.PubkeyToAddress(private.PublicKey)
	public := private.PublicKey
	wallet := Wallet{private, public}

	return &wallet
}

// GetAddress returns wallet address
func (w Wallet) GetAddress() common.Address {
	address := crypto.PubkeyToAddress(w.PublicKey)
	return address
}

// // ValidateAddress check if address if valid
// func ValidateAddress(address common.Address) bool {
// 	pubKeyHash := Base58Decode([]byte(address))
// 	actualChecksum := pubKeyHash[len(pubKeyHash)-addressChecksumLen:]
// 	version := pubKeyHash[0]
// 	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-addressChecksumLen]
// 	targetChecksum := checksum(append([]byte{version}, pubKeyHash...))

// 	return bytes.Compare(actualChecksum, targetChecksum) == 0
// }
