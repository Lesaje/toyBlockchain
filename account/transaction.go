package account

// Should CreateTransaction be called from Mempool?

import (
    "crypto/sha256"
    "errors"
    "fmt"
    "hash"
)

type Transaction struct {
    TransactionID   hash.Hash
    SetOfOperations []Operation
}

func (bch Blockchain) CreateTransaction(SetOfOperations []Operation) (Transaction, error) {
    var output Transaction
    var bytes []byte

    for _, op := range SetOfOperations {
        if bch.Verify(op) != true {
            return output, errors.New("incorrect operation")
        }
        bytes = append(bytes, op.ToByte()...)
    }

    txHash := sha256.New()
    txHash.Write(bytes)

    output.TransactionID = txHash
    output.SetOfOperations = SetOfOperations

    return output, nil
}

func (tx Transaction) ToByte() ([]byte, error) {
    var output []byte
    output = append(output, tx.TransactionID.Sum(nil)...)

    for _, el := range tx.SetOfOperations {
        output = append(output, el.ToByte()...)
    }

    return output, nil
}

func (tx Transaction) String() string {
    var str string
    str += fmt.Sprintf("TransactionID: %x", tx.TransactionID.Sum(nil)) + "\n"
    for _, el := range tx.SetOfOperations {
        str += el.String() + "\n"
    }
    return str
}
