# Blockchain

[![GoDoc](https://godoc.org/github.com/lynn9388/blockchain?status.svg)](https://godoc.org/github.com/lynn9388/blockchain)
[![Build Status](https://travis-ci.com/lynn9388/blockchain.svg?branch=master)](https://travis-ci.com/lynn9388/blockchain)

A simple blockchain implementation.

## Install

Fist, use `go get` to install the latest version of the library:

```sh
go get -u github.com/lynn9388/blockchain
```

Next, include SupSub in your application:

```go
import "github.com/lynn9388/blockchain"
```

## Example

```go
genesis := GetGenesisBlock()
bc := NewBlockchain(genesis)
bc.Add(genesis.NewBlock([]byte("lynn9388")))
```