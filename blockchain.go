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
	"errors"
	"strconv"
)

// Blockchain records all the blocks.
type Blockchain []Block

// NewBlockchain returns a blockchain witch records only genesis block.
func NewBlockchain() Blockchain {
	return []Block{*GetGenesisBlock()}
}

// GetPrevBlock return a previous block based on the index.
func (bc *Blockchain) GetPrevBlock(b *Block) (*Block, error) {
	index := b.Index - 1
	if index > len(*bc) {
		return nil, errors.New("index out or range: " + strconv.Itoa(index))
	}
	return &((*bc)[index]), nil
}

// Add append a block to the blockchain if the block is valid.
func (bc *Blockchain) Add(b *Block) error {
	prev, err := bc.GetPrevBlock(b)
	if err != nil {
		return err
	}

	if !b.isValid(prev) {
		return errors.New("block is not valid")
	}

	*bc = append(*bc, *b)
	return nil
}
