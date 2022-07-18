package account

import (
    "crypto/sha256"
    "fmt"
    "hash"
)

type Block struct {
    BlockId      hash.Hash
    PrevHash     hash.Hash
    Transactions []Transaction
}

func CreateBlock(prevHash hash.Hash, transactions []Transaction) (Block, error) {
    blockBytes := make([]byte, 0)
    //var blockBytes []byte
    if prevHash != nil {
        blockBytes = append(blockBytes, prevHash.Sum(nil)...)
    }
    var output Block

    for _, el := range transactions {
        txBytes, err := el.ToByte()
        if err != nil {
            return output, err
        }
        blockBytes = append(blockBytes, txBytes...)
    }

    blockHash := sha256.New()
    blockHash.Write(blockBytes)

    output.BlockId = blockHash
    output.PrevHash = prevHash
    output.Transactions = transactions
    return output, nil
}

func (b Block) String() string {
    var str string
    str += fmt.Sprintf("BlockId: %x\n", b.BlockId.Sum(nil))
    if b.PrevHash == nil {
        str += fmt.Sprintf("PrevHash: %x\n", 0)
    } else {
        str += fmt.Sprintf("PrevHash: %x\n", b.PrevHash.Sum(nil))
    }
    for _, el := range b.Transactions {
        str += el.String() + "\n"
    }
    return str
}
