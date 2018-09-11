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
	"encoding/hex"
	"github.com/lynn9388/merkletree"
	"strconv"
	"time"
)

// BlockHeader holds the metadata of a block
type BlockHeader struct {
	Index      int    `json:"index"`
	Time       int64  `json:"time"`
	PrevHash   string `json:"prevHash"`
	MerkleRoot string `json:"merkleRoot"`
}

// Block holds batches of valid transactions.
type Block struct {
	BlockHeader
	Data []merkletree.Data `json:"data"`
}

// ToByte converts the block to a slice of byte.
func (b *Block) ToByte() []byte {
	var buff bytes.Buffer
	buff.WriteString(strconv.Itoa(b.Index))
	buff.WriteString(strconv.FormatInt(b.Time, 10))
	buff.Write([]byte(b.PrevHash))
	buff.Write([]byte(b.MerkleRoot))
	for _, datum := range b.Data {
		buff.Write(datum.ToByte())
	}
	return buff.Bytes()
}

// Hash returns the SHA256 hash values in hexadecimal of the data.
func (b *Block) Hash() string {
	hash := sha256.Sum256(b.ToByte())
	return hex.EncodeToString(hash[:])
}

// NewBlock creates a new block next to current block.
func (b *Block) NewBlock(data ...merkletree.Data) *Block {
	return &Block{
		BlockHeader: BlockHeader{
			Index:      b.Index + 1,
			Time:       time.Now().Unix(),
			PrevHash:   b.Hash(),
			MerkleRoot: merkletree.NewMerkleTree(data...).Root.Hash,
		},
		Data: data,
	}
}

// isValid checks if every fields in a block is valid.
func (b *Block) isValid(prevBlock *Block) bool {
	if b.Index != prevBlock.Index+1 ||
		b.Time < prevBlock.Time ||
		b.PrevHash != prevBlock.Hash() ||
		b.MerkleRoot != merkletree.NewMerkleTree(b.Data...).Root.Hash {
		return false
	}
	return true
}

// GenesisBlock returns the genesis block.
func GenesisBlock() *Block {
	t, _ := time.Parse("2006-1-02", "1993-8-08")
	return &Block{
		BlockHeader: BlockHeader{
			Index:      0,
			Time:       t.Unix(),
			PrevHash:   "",
			MerkleRoot: "",
		},
		Data: nil,
	}
}
