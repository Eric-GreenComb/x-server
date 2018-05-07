package ether

import (
	"crypto/ecdsa"
	crand "crypto/rand"
	"io"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pborman/uuid"
)

const (
	keyHeaderKDF = "scrypt"

	// StandardScryptN is the N parameter of Scrypt encryption algorithm, using 256MB
	// memory and taking approximately 1s CPU time on a modern processor.
	StandardScryptN = 1 << 18

	// StandardScryptP is the P parameter of Scrypt encryption algorithm, using 256MB
	// memory and taking approximately 1s CPU time on a modern processor.
	StandardScryptP = 1

	// LightScryptN is the N parameter of Scrypt encryption algorithm, using 4MB
	// memory and taking approximately 100ms CPU time on a modern processor.
	LightScryptN = 1 << 12

	// LightScryptP is the P parameter of Scrypt encryption algorithm, using 4MB
	// memory and taking approximately 100ms CPU time on a modern processor.
	LightScryptP = 6

	scryptR     = 8
	scryptDKLen = 32

	// VeryLightScryptN VeryLightScryptN
	VeryLightScryptN = 2
	// VeryLightScryptP VeryLightScryptP
	VeryLightScryptP = 1
)

// Ks Keystore
var Ks Keystore

// Keystore Keystore
type Keystore struct {
}

// NewKey NewKey
func (ks *Keystore) NewKey() (*keystore.Key, error) {
	key, err := ks.newKey(crand.Reader)
	return key, err
}

// GenKeystore GenKeystore
func (ks *Keystore) GenKeystore(key *keystore.Key, passphrase string) ([]byte, error) {
	keyjson, err := keystore.EncryptKey(key, passphrase, VeryLightScryptN, VeryLightScryptP)
	return keyjson, err
}

// Update Update
func (ks *Keystore) Update(keyjson []byte, passphrase, newPassphrase string) ([]byte, error) {
	key, err := keystore.DecryptKey(keyjson, passphrase)
	if err != nil {
		return nil, err
	}

	_keyjson, err := keystore.EncryptKey(key, newPassphrase, VeryLightScryptN, VeryLightScryptP)
	return _keyjson, err
}

func (ks *Keystore) newKey(rand io.Reader) (*keystore.Key, error) {
	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand)
	if err != nil {
		return nil, err
	}
	return ks.newKeyFromECDSA(privateKeyECDSA), nil
}

func (ks *Keystore) newKeyFromECDSA(privateKeyECDSA *ecdsa.PrivateKey) *keystore.Key {
	id := uuid.NewRandom()
	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}
	return key
}
