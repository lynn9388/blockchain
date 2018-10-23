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
	"encoding/binary"
	"time"

	"github.com/lynn9388/merkletree"
	"go.uber.org/zap"
)

// BlockHeader holds the metadata of a block
type BlockHeader struct {
	Index      int    `json:"index"`
	Time       int64  `json:"time"`
	PrevHash   []byte `json:"prevHash"`
	MerkleRoot []byte `json:"merkleRoot"`
	Extra      []byte `json:"extra"`
}

// Block holds batches of valid data/transactions.
type Block struct {
	Header *BlockHeader `json:"header"`
	Data   [][]byte     `json:"data"`
}

var log *zap.SugaredLogger

func init() {
	logger, _ := zap.NewDevelopment()
	log = logger.Sugar()
}

// ToByte converts the block header to bytes.
func (bh *BlockHeader) ToByte() []byte {
	var buff bytes.Buffer
	buff.Write(intToByte(int64(bh.Index)))
	buff.Write(intToByte(bh.Time))
	buff.Write(bh.PrevHash)
	buff.Write(bh.MerkleRoot)
	buff.Write(bh.Extra)
	return buff.Bytes()
}

// Hash returns the SHA256 hash of the block header.
func (bh *BlockHeader) Hash() []byte {
	hash := sha256.Sum256(bh.ToByte())
	return hash[:]
}

// IsValid checks if every fields in a block is valid. isExtraValid can be
// nil if check extra info is unnecessary.
func (b *Block) IsValid(prevBlockHeader *BlockHeader, isExtraValid func([]byte) bool) bool {
	if b.Header.Index != prevBlockHeader.Index+1 ||
		b.Header.Time < prevBlockHeader.Time ||
		!bytes.Equal(b.Header.PrevHash, prevBlockHeader.Hash()) ||
		!bytes.Equal(b.Header.MerkleRoot, merkletree.NewMerkleTree(b.Data...).Root.Hash) ||
		isExtraValid != nil && !isExtraValid(b.Header.Extra) {
		return false
	}
	return true
}

// NewGenesisBlock returns the genesis block.
func NewGenesisBlock() *Block {
	t, err := time.Parse("2006-1-02", "1993-8-08")
	if err != nil {
		log.Panic(err)
	}

	return &Block{
		Header: &BlockHeader{
			Index:      0,
			Time:       t.Unix(),
			PrevHash:   []byte(""),
			MerkleRoot: []byte(""),
			Extra:      []byte(""),
		},
		Data: [][]byte{},
	}
}

// NewBlock creates a new block.
func NewBlock(prevBlockHeader *BlockHeader, extra []byte, data [][]byte) *Block {
	return &Block{
		Header: &BlockHeader{
			Index:      prevBlockHeader.Index + 1,
			Time:       time.Now().Unix(),
			PrevHash:   prevBlockHeader.Hash(),
			MerkleRoot: merkletree.NewMerkleTree(data...).Root.Hash,
			Extra:      extra,
		},
		Data: data,
	}
}

// intToByte converts an int64 to a byte slice.
func intToByte(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
