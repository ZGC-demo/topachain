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
package types

import (
	"bytes"
	"encoding/gob"

	"github.com/mintzhao/topachain/common/crypto"
)

// Transaction
type Transaction struct {
	Payload []byte // The data came from Blockchain Application, it's transparent to the ToPa Chain.
}

func (tx *Transaction) Hash() ([]byte, error) {
	return crypto.Hash(tx.Payload)
}

func (tx *Transaction) Equals(other interface{}) bool {
	othertx, ok := other.(*Transaction)
	if !ok {
		return false
	}

	return bytes.Equal(tx.Payload, othertx.Payload)
}

func init() {
	gob.Register(Transaction{})
}
