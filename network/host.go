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

// Package network is a simple blockchain network implementation.
package network

import (
	"github.com/lynn9388/blockchain"
	"github.com/lynn9388/p2p"
	"go.uber.org/zap"
)

// Host is a full node in blockchain network.
type Host struct {
	*p2p.Node
	Blockchain *blockchain.Blockchain
}

var log *zap.SugaredLogger

func init() {
	logger, _ := zap.NewDevelopment()
	log = logger.Sugar()
}

// NewHost creates a new initialized host.
func NewHost(addr string) *Host {
	return &Host{
		Node:       p2p.NewNode(addr),
		Blockchain: blockchain.NewBlockchain("addr"),
	}
}

// TODO blockchain synchronization
