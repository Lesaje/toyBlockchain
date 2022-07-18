package main

import (
    a "blokchain/account"
    "blokchain/keypair"
    "fmt"
    "testing"
)

func TestBlockchain(t *testing.T) {
    newBlockchain, genesisAcc, err := a.InitBlockchain(1000000)
    if err != nil {
        t.Fatalf("Error initializing blockchain: %v", err)
    }
    accSlice := make([]a.Account, 10)
    for i := 0; i < 10; i++ {
        var newAccKeyPair keypair.KeyPair
        err := newAccKeyPair.GenerateKeyPair()
        if err != nil {
            t.Fatalf("Error generating keypair: %v", err)
        }
        accSlice[i] = newBlockchain.NewAccount(newAccKeyPair)
    }

    operationSet := make([]a.Operation, 10)
    for i, el := range accSlice {
        operationSet[i], err = genesisAcc.CreateOp(el.Identifier, 10)
        if err != nil {
            t.Fatalf("Error creating operation: %v", err)
        }
    }
    tx, err := newBlockchain.CreateTransaction(operationSet)
    if err != nil {
        t.Fatalf("Error creating transaction: %v", err)
    }

    block, err := a.CreateBlock(newBlockchain.BlockHistory[0].BlockId, []a.Transaction{tx})
    if err != nil {
        t.Fatalf("Error creating block: %v", err)
    }

    err = newBlockchain.AddBlock(block)
    if err != nil {
        t.Fatalf("Error adding block: %v", err)
    }

    fmt.Println(newBlockchain)
}
