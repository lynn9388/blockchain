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
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	bc := NewBlockchain("test")

	if bc.DB == nil || len(bc.Tips) != 1 || bc.BestTip == nil || bc.Tips[0] != bc.BestTip {
		t.Errorf("%+v", bc)
	}

	bc.DB.Close()
	bc = NewBlockchain("test")
	if bc.DB == nil || len(bc.Tips) != 1 || bc.BestTip == nil || bc.Tips[0] != bc.BestTip {
		t.Errorf("%+v", bc)
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
