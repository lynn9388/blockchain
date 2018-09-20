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
	"fmt"
	"strconv"
)

// Blockchain records all the blocks.
type Blockchain struct {
	Blocks []*Block
}

// NewBlockchain returns a blockchain which records only genesis block.
func NewBlockchain(genesis *Block) *Blockchain {
	return &Blockchain{
		Blocks: []*Block{genesis},
	}
}

// AddBlock appends a block to the blockchain if the block is valid.
func (bc *Blockchain) AddBlock(b *Block) error {
	if b.Header.Index != bc.Length() {
		return errors.New("failed to add block to the end")
	}

	prev, err := bc.GetBlock(b.Header.Index - 1)
	if err != nil {
		return err
	}

	if !b.IsValid(&prev.Header) {
		return errors.New("block is not valid")
	}

	bc.Blocks = append(bc.Blocks, b)
	return nil
}

// GetBlock returns a block based on its index.
func (bc *Blockchain) GetBlock(i int) (*Block, error) {
	if i < 0 || i > bc.Length() {
		return nil, fmt.Errorf("index out of range: %v", strconv.Itoa(i))
	}
	return bc.Blocks[i], nil
}

// Length returns the length of the blockchain.
func (bc *Blockchain) Length() int {
	return len(bc.Blocks)
}
