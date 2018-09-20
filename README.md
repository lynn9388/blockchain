# Blockchain

[![GoDoc](https://godoc.org/github.com/lynn9388/blockchain?status.svg)](https://godoc.org/github.com/lynn9388/blockchain)
[![Build Status](https://travis-ci.com/lynn9388/blockchain.svg?branch=master)](https://travis-ci.com/lynn9388/blockchain)

A simple blockchain implementation.

## Install

Fist, use `go get` to install the latest version of the library:

```sh
go get -u github.com/lynn9388/blockchain
```

Next, include this package in your application:

```go
import "github.com/lynn9388/blockchain"
```

## Example

```go
genesis := GenesisBlock()
bc := NewBlockchain(genesis)
bc.AddBlock(genesis.Header.NewBlock(StringData("lynn9388")))
```