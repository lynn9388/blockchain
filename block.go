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

// Package blockchain provides a simple blockchain implementation.
package blockchain

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block holds batches of valid transactions.
type Block struct {
	Index    int    `json:"index"`
	Time     int64  `json:"time"`
	PrevHash []byte `json:"prevHash"`
	Data     []byte `json:"data"`
}

// ToByte convert the block to a slice of byte.
func (b *Block) ToByte() []byte {
	var buff bytes.Buffer
	buff.WriteString(strconv.Itoa(b.Index))
	buff.WriteString(strconv.FormatInt(b.Time, 10))
	buff.Write(b.PrevHash)
	buff.Write(b.Data)
	return buff.Bytes()
}

// Hash returns the SHA256 hash values of the block.
func (b *Block) Hash() []byte {
	h := sha256.New()
	h.Write(b.ToByte())
	return h.Sum(nil)
}

// NewBlock creates a new block.
func (b *Block) NewBlock(prevBlock *Block, data []byte) *Block {
	return &Block{
		Index:    prevBlock.Index + 1,
		Time:     time.Now().Unix(),
		PrevHash: prevBlock.Hash(),
		Data:     data,
	}
}

// isValid checks if every fields in a block is valid.
func (b *Block) isValid(prevBlock *Block) bool {
	if b.Index != prevBlock.Index+1 ||
		b.Time <= prevBlock.Time ||
		bytes.Compare(b.PrevHash, prevBlock.Hash()) != 0 {
		return false
	}
	return true
}

// GetGenesisBlock returns the genesis block.
func GetGenesisBlock() *Block {
	t, _ := time.Parse("2006-1-02", "1993-8-08")
	return &Block{
		Index:    0,
		Time:     t.Unix(),
		PrevHash: []byte{},
		Data:     []byte{},
	}
}
