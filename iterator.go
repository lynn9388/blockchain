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
	"github.com/boltdb/bolt"
)

// Iterator is used to iterate over blockchain's blocks.
type Iterator struct {
	db        *bolt.DB
	blockHash []byte
}

// NewIterator returns a iterator with a start block.
func NewIterator(bc *Blockchain, blockHash []byte) *Iterator {
	return &Iterator{db: bc.DB, blockHash: blockHash}
}

// Prev returns the block pointed by the iterator and move to the previous
// block. It will end with genesis block and then return nil.
func (i *Iterator) Prev() *Block {
	var block Block

	if len(i.blockHash) == 0 {
		return nil
	}

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		get(b, i.blockHash, &block)
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	i.blockHash = block.Header.PrevHash
	return &block
}
