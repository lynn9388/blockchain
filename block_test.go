/*
 * Copyright © 2018 Lynn <lynn9388@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package blockchain

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func TestBlockHeader_ToByte(t *testing.T) {
	genesis := NewGenesisBlock()
	if hex.EncodeToString(genesis.Header.ToByte()) != "0000000000000000000000002c644200" {
		t.FailNow()
	}
}

func TestBlockHeader_Hash(t *testing.T) {
	genesis := NewGenesisBlock()
	if hex.EncodeToString(genesis.Header.Hash()) != "6d46d7f95388521ac55906cffda76f024b8fc54f0537ee753f373b3a85b2a1dc" {
		t.FailNow()
	}
}

func TestBlock_IsValid(t *testing.T) {
	genesis := NewGenesisBlock()
	block := NewBlock(genesis.Header, []byte(""), [][]byte{[]byte("lynn9388")})
	if block.IsValid(genesis.Header, nil) == false {
		t.FailNow()
	}

	if block.IsValid(genesis.Header, func(b []byte) bool {
		return bytes.Equal(b, genesis.Header.Extra)
	}) == false {
		t.FailNow()
	}

	if block.IsValid(genesis.Header, func(b []byte) bool {
		return !bytes.Equal(b, genesis.Header.Extra)
	}) == true {
		t.FailNow()
	}

	block.Data = [][]byte{[]byte("9388lynn")}
	if block.IsValid(genesis.Header, nil) == true {
		t.FailNow()
	}
}
