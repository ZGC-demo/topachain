// Copyright © 2018 Zhao Ming <mint.zhao.chiu@gmail.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package crypto

import (
	"crypto"
	"sync"

	"github.com/mintzhao/topachain/common/crypto/hasher"
	"github.com/mintzhao/topachain/common/crypto/signer"
	_ "github.com/mintzhao/topachain/common/logging"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
)

var logger = logging.MustGetLogger("crypto")

// CryptoInterface contains all the functions related to crypto, include hash, encrypt, sign etc.
type CryptoInterface interface {
	hasher.Hasher
	signer.Signer
}

// cryptoImpl implement CryptoInterface
type cryptoImpl struct {
	her  hasher.Hasher
	sger signer.Signer
}

// NewCrypto return a CryptoInterface instance based on configuration
func NewCrypto() (CryptoInterface, error) {
	hashName := getHasherName()
	her, err := hasher.GetHasher(hashName)
	if err != nil {
		return nil, err
	}

	signerName := getSignerName()
	sger, err := signer.GetSigner(signerName)
	if err != nil {
		return nil, err
	}

	return &cryptoImpl{
		her:  her,
		sger: sger,
	}, nil
}

// Hash is a global function to hashes message msg.
func Hash(msg []byte, hashName ...string) ([]byte, error) {
	if len(hashName) != 0 {
		her, err := hasher.GetHasher(hashName[0])
		if err != nil {
			return nil, err
		}

		return her.Hash(msg)
	}

	return getInstance().Hash(msg)
}

// Hash hashes messages msg
func (impl *cryptoImpl) Hash(msg []byte) ([]byte, error) {
	return impl.her.Hash(msg)
}

// Sign is a global function to signs msg using PrivateKey k.
func Sign(k crypto.PrivateKey, msg []byte, opts crypto.SignerOpts, signerName ...string) ([]byte, error) {
	if len(signerName) != 0 {
		sger, err := signer.GetSigner(signerName[0])
		if err != nil {
			return nil, err
		}

		return sger.Sign(k, msg, opts)
	}

	return getInstance().Sign(k, msg, opts)
}

// Sign signs digest using PrivateKey k.
func (impl *cryptoImpl) Sign(k crypto.PrivateKey, msg []byte, opts crypto.SignerOpts) ([]byte, error) {
	return impl.sger.Sign(k, msg, opts)
}

// Verify is a global function to verify signature
func Verify(k crypto.PublicKey, signature, msg []byte, opts crypto.SignerOpts, signerName ...string) (bool, error) {
	if len(signerName) != 0 {
		sger, err := signer.GetSigner(signerName[0])
		if err != nil {
			return false, err
		}

		return sger.Verify(k, signature, msg, opts)
	}

	return getInstance().Verify(k, signature, msg, opts)
}

// Verify verifies signature against key k and digest
func (impl *cryptoImpl) Verify(k crypto.PublicKey, signature, msg []byte, opts crypto.SignerOpts) (bool, error) {
	return impl.sger.Verify(k, signature, msg, opts)
}

var (
	instance CryptoInterface
	once     sync.Once
)

func getInstance() CryptoInterface {
	once.Do(func() {
		var err error
		instance, err = NewCrypto()
		if err != nil {
			logger.Panic(err)
		}
	})

	return instance
}

// getHasherName returns customize hash function, default is SHA256
func getHasherName() string {
	hasherName := viper.GetString("common.crypto.hash")

	if hasherName == "" {
		hasherName = "SHA256"
	}

	return hasherName
}

// getSignerName returns customize signer, default is ECDSA
func getSignerName() string {
	signerName := viper.GetString("common.crypto.sign")

	if signerName == "" {
		signerName = "ECDSA"
	}

	return signerName
}
