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
	"strconv"
	"time"

	"github.com/lynn9388/merkletree"
)

// BlockHeader holds the metadata of a block
type BlockHeader struct {
	Index      int    `json:"index"`
	Time       int64  `json:"time"`
	PrevHash   string `json:"prevHash"`
	MerkleRoot string `json:"merkleRoot"`
}

// Block holds batches of valid data/transactions.
type Block struct {
	Header BlockHeader
	Data   []Data `json:"data"`
}

// ToByte converts the block header to bytes.
func (bh *BlockHeader) ToByte() []byte {
	var buff bytes.Buffer
	buff.WriteString(strconv.Itoa(bh.Index))
	buff.WriteString(strconv.FormatInt(bh.Time, 10))
	buff.WriteString(bh.PrevHash)
	buff.WriteString(bh.MerkleRoot)
	return buff.Bytes()
}

// Hash returns the SHA256 hash values in hexadecimal of the block header.
func (bh *BlockHeader) Hash() string {
	hash := sha256.Sum256(bh.ToByte())
	return hex.EncodeToString(hash[:])
}

// NewBlock creates a new block next to current block header.
func (bh *BlockHeader) NewBlock(data ...Data) *Block {
	var db [][]byte
	for _, datum := range data {
		db = append(db, datum.ToByte())
	}

	return &Block{
		Header: BlockHeader{
			Index:      bh.Index + 1,
			Time:       time.Now().Unix(),
			PrevHash:   bh.Hash(),
			MerkleRoot: merkletree.NewMerkleTree(db...).Root.Hash,
		},
		Data: data,
	}
}

// IsValid checks if every fields in a block is valid.
func (b *Block) IsValid(prevBlockHeader *BlockHeader) bool {
	var db [][]byte
	for _, datum := range b.Data {
		db = append(db, datum.ToByte())
	}

	if b.Header.Index != prevBlockHeader.Index+1 ||
		b.Header.Time < prevBlockHeader.Time ||
		b.Header.PrevHash != prevBlockHeader.Hash() ||
		b.Header.MerkleRoot != merkletree.NewMerkleTree(db...).Root.Hash {
		return false
	}
	return true
}

// GenesisBlock returns the genesis block.
func GenesisBlock() *Block {
	t, _ := time.Parse("2006-1-02", "1993-8-08")
	return &Block{
		Header: BlockHeader{
			Index:      0,
			Time:       t.Unix(),
			PrevHash:   "",
			MerkleRoot: "",
		},
		Data: make([]Data, 0),
	}
}
