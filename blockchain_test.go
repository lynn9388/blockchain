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
	"testing"
)

func TestBlockchain_AddBlock(t *testing.T) {
	genesis := NewGenesisBlock()
	bc := NewBlockchain(genesis)

	err := bc.AddBlock(genesis.Header.NewBlock(StringData("lynn9388")))
	if err != nil {
		t.Error(err)
	}
	if bc.Length() != 2 {
		t.Errorf("failed to add block: %v", bc.Length())
	}

	err = bc.AddBlock(genesis.Header.NewBlock(StringData("lynn9388")))
	if err == nil || bc.Length() != 2 {
		t.Error("failed to add duplicate block")
	}

	index := bc.Length() - 1
	b, err := bc.GetBlock(index)
	if err != nil {
		t.Errorf("failed to get block: %v", index)
	}
	err = bc.AddBlock(b.Header.NewBlock(StringData("lynn9388")))
	if err != nil {
		t.Error(err)
	}
	if bc.Length() != 3 {
		t.Errorf("failed to add block: %v", bc.Length())
	}
}
