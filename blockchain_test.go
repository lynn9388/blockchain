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
	"os"
	"testing"

	"github.com/boltdb/bolt"
)

func TestNewBlockchain(t *testing.T) {
	bc := NewBlockchain("test.db")

	if bc.DB == nil || len(bc.Tips) != 1 || bc.BestTip == nil || bc.Tips[0] != bc.BestTip {
		t.Errorf("%+v", bc)
	}

	bc.DB.Close()
	bc = NewBlockchain("test.db")
	defer os.Remove("test.db")
	defer bc.DB.Close()
	if bc.DB == nil || len(bc.Tips) != 1 || bc.BestTip == nil || bc.Tips[0] != bc.BestTip {
		t.Errorf("%+v", bc)
	}
}

func TestBlockchain_AddBlock(t *testing.T) {
	bc := NewBlockchain("test.db")

	tests := []string{"lynn", "9388"}
	var hash []byte
	for _, test := range tests {
		block := NewBlock(bc.BestTip.Header, nil, [][]byte{[]byte(test)})
		bc.AddBlock(block, nil)

		isInTips := false
		for _, tip := range bc.Tips {
			if tip == block {
				isInTips = true
				break
			}
		}
		if !isInTips {
			t.Error("failed to find block in tips")
		}

		if bc.BestTip != block {
			t.Error("failed to check bestTip")
		}

		hash = block.Header.Hash()
	}

	bc.DB.Close()
	bc = NewBlockchain("test.db")
	defer os.Remove("test.db")
	defer bc.DB.Close()
	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		var bestTip []byte
		get(b, []byte(bestTipKey), &bestTip)
		if hash == nil || !bytes.Equal(hash, bestTip) {
			t.Error("failed to check bestTip in DB")
		}
		return nil
	})
	if err != nil {
		t.FailNow()
	}
}

func TestSerialize(t *testing.T) {
	tests := []interface{}{"lynn9388", NewGenesisBlock()}
	for test := range tests {
		b := serialize(test)
		if b == nil || len(b) == 0 {
			t.Fail()
		}
	}
}

func TestDeserialize(t *testing.T) {
	testStr := "lynn9388"
	b := serialize(testStr)
	var str string
	deserialize(b, &str)
	if str != testStr {
		t.Error("failed to deserialize string")
	}

	testBlock := NewGenesisBlock()
	b = serialize(testBlock)
	var block Block
	deserialize(b, &block)
	if !bytes.Equal(block.Header.Hash(), testBlock.Header.Hash()) {
		t.Error("failed to deserialize block")
	}
}
