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

import "testing"

type testData string

func (td testData) ToByte() []byte {
	return []byte(td)
}

func TestBlockchain_Add(t *testing.T) {
	genesis := GenesisBlock()
	bc := NewBlockchain(genesis)
	b := genesis.NewBlock(testData("lynn9388"))
	err := bc.Add(b)
	if err != nil || len(bc) != 2 {
		t.Error(err)
	}

	err = bc.Add(genesis.NewBlock(testData("lynn9388")))
	if err == nil || len(bc) != 2 {
		t.Fail()
	}

	err = bc.Add(b.NewBlock(testData("lynn9388")))
	if err != nil || len(bc) != 3 {
		t.Error(err)
	}
}
