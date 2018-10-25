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
	"bytes"
	"encoding/gob"
	"fmt"
	"os"

	"github.com/boltdb/bolt"
)

const (
	dbFile       = "blockchain_%s.db" // name of DB file
	blocksBucket = "blocks"           // name of bucket storing blockchain
	tipsKey      = "tips"             // key for last block's hash of all branches
	bestTipKey   = "bestTip"          // key for last block's hash of longest branch
)

// Blockchain implements interactions with a DB.
type Blockchain struct {
	DB      *bolt.DB // DB stored the blockchain
	Tips    []*Block // last block of all branches
	BestTip *Block   // last block of longest branch
}

// NewBlockchain creates a blockchain from DB file. If the file does not
// exist then it will be created and initialize a new blockchain with
// genesis block.
func NewBlockchain(nodeID string) *Blockchain {
	var db *bolt.DB
	var tips []*Block
	var bestTip *Block

	dbFile := fmt.Sprintf(dbFile, nodeID)
	var err error
	if !dbExists(dbFile) {
		db, err = bolt.Open(dbFile, 0600, nil)
		if err != nil {
			log.Panic(err)
		}

		err = db.Update(func(tx *bolt.Tx) error {
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Panic(err)
			}

			genesis := NewGenesisBlock()
			tips = append(tips, genesis)
			bestTip = tips[0]

			put(b, genesis.Header.Hash(), genesis)
			putTips(b, tips)
			put(b, []byte(bestTipKey), bestTip.Header.Hash())

			return nil
		})
		if err != nil {
			log.Panic()
		}
	} else {
		db, err = bolt.Open(dbFile, 0600, nil)
		if err != nil {
			log.Panic(err)
		}

		err = db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blocksBucket))
			tips = getTips(b)

			var hash []byte
			get(b, []byte(bestTipKey), &hash)
			for _, tip := range tips {
				if bytes.Equal(tip.Header.Hash(), hash) {
					bestTip = tip
				}
			}

			return nil
		})
		if err != nil {
			log.Panic(err)
		}
	}

	return &Blockchain{DB: db, Tips: tips, BestTip: bestTip}
}

// AddBlock saves a block into the blockchain if it is valid. isExtraValid
// can be nil if check extra info is unnecessary.
func (bc *Blockchain) AddBlock(b *Block, isExtraValid func([]byte) bool) {
	tipNum := -1
	for i, tip := range bc.Tips {
		if bytes.Equal(tip.Header.Hash(), b.Header.PrevHash) {
			tipNum = i
			break
		}
	}

	if tipNum < 0 {
		log.Panic("failed to find previous block")
	}

	if !b.IsValid(bc.Tips[tipNum].Header, isExtraValid) {
		log.Panic("block is not valid")
	}

	err := bc.DB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))

		put(bucket, b.Header.Hash(), *b)

		bc.Tips[tipNum] = b
		putTips(bucket, bc.Tips)

		if b.Header.Index > bc.BestTip.Header.Index {
			bc.BestTip = b
			put(bucket, []byte(bestTipKey), b.Header.Hash())
		}

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

// dbExists checks if DB file exists.
func dbExists(dbFile string) bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}

// serialize serializes data to byte slice.
func serialize(data interface{}) []byte {
	var buff bytes.Buffer
	encoder := gob.NewEncoder(&buff)
	err := encoder.Encode(data)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}

// deserialize deserializes serialized byte slice back to data.
func deserialize(b []byte, dataPtr interface{}) {
	decoder := gob.NewDecoder(bytes.NewReader(b))
	err := decoder.Decode(dataPtr)
	if err != nil {
		log.Panic(err)
	}
}

// put sets the value for a key in the bucket. The value will be serialized
// automatically.
func put(b *bolt.Bucket, key []byte, value interface{}) {
	err := b.Put(key, serialize(value))
	if err != nil {
		log.Panic(fmt.Sprintf("failed to put %v", key), err)
	}
}

// get retrieves the value for a key in the bucket, and deserializes the
// value to dataPtr.
func get(b *bolt.Bucket, key []byte, dataPtr interface{}) {
	data := b.Get(key)
	if data == nil {
		log.Panic(fmt.Sprintf("failed to get %v", key))
	}
	deserialize(data, dataPtr)
}

// putTips puts the hashes of tip blocks in the bucket.
func putTips(b *bolt.Bucket, tips []*Block) {
	var hashes [][]byte
	for _, tip := range tips {
		hashes = append(hashes, tip.Header.Hash())
	}
	put(b, []byte(tipsKey), hashes)
}

// getTips retrieves the tip blocks from the bucket.
func getTips(b *bolt.Bucket) []*Block {
	var hashes [][]byte
	var tips []*Block

	get(b, []byte(tipsKey), &hashes)
	for _, hash := range hashes {
		var tip Block
		get(b, hash, &tip)
		tips = append(tips, &tip)
	}

	return tips
}
