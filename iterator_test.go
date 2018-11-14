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
	"os"
	"testing"
)

func TestIterator_Prev(t *testing.T) {
	bc := NewBlockchain("test.db", NewGenesisBlock())
	defer os.Remove("test.db")
	defer bc.DB.Close()

	tests := []string{"lynn", "9388"}
	for _, test := range tests {
		block := NewBlock(bc.BestTip.Header, nil, [][]byte{[]byte(test)})
		bc.AddBlock(block, nil)
	}

	iterator := NewIterator(bc, bc.BestTip.Header.Hash())
	count := 0
	for {
		block := iterator.Prev()
		if block == nil {
			break
		}
		count++
	}

	if count != len(tests)+1 {
		t.Errorf("iterator iterates %v blocks", count)
	}

}
