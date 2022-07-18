package account

import (
    "blokchain/keypair"
    "errors"
    "fmt"
)

type Blockchain struct {
    CoinDatabase map[string]uint64
    BlockHistory []Block
    Coins        uint64
}

func InitBlockchain(Coins uint64) (Blockchain, Account, error) {
    var output Blockchain
    output.CoinDatabase = make(map[string]uint64)

    var depositoryKeyPair keypair.KeyPair
    err := depositoryKeyPair.GenerateKeyPair()
    if err != nil {
        return output, Account{}, err
    }
    genesisDepository := NewDepository(depositoryKeyPair, Coins)
    depositoryId := fmt.Sprintf("%x", genesisDepository.Identifier)
    output.CoinDatabase[depositoryId] = Coins

    var genesisKeyPair keypair.KeyPair
    err = genesisKeyPair.GenerateKeyPair()
    if err != nil {
        return output, Account{}, err
    }
    genesisAccount := output.NewAccount(genesisKeyPair)
    genesisId := fmt.Sprintf("%x", genesisAccount.Identifier)
    output.CoinDatabase[genesisId] = 0

    //send 0.01% of all coins to the first account
    operation, err := genesisDepository.SendInterests(genesisAccount.Identifier, Coins/10000)
    if err != nil {
        return output, Account{}, err
    }

    setOfOperations := make([]Operation, 1)
    setOfOperations[0] = operation
    tx, err := output.CreateTransaction(setOfOperations)
    if err != nil {
        return output, Account{}, err
    }

    setOfTransactions := make([]Transaction, 1)
    setOfTransactions[0] = tx
    block, err := CreateBlock(nil, setOfTransactions)
    if err != nil {
        return output, Account{}, err
    }

    err = output.AddBlock(block)
    if err != nil {
        return output, Account{}, err
    }
    output.Coins = Coins
    return output, genesisAccount, nil
}

func (bch *Blockchain) AddBlock(b Block) error {
    if bch.validateBlock(b) == false {
        return errors.New("invalid block")
    }

    for _, tx := range b.Transactions {
        for _, op := range tx.SetOfOperations {
            sender := fmt.Sprintf("%x", op.Sender)
            receiver := fmt.Sprintf("%x", op.Receiver)
            bch.CoinDatabase[sender] -= op.Amount
            bch.CoinDatabase[receiver] += op.Amount
        }
    }
    bch.BlockHistory = append(bch.BlockHistory, b)
    return nil
}

func (bch Blockchain) validateBlock(b Block) bool {

    if len(bch.BlockHistory) != 0 {
        if b.PrevHash != bch.BlockHistory[len(bch.BlockHistory)-1].BlockId {
            return false
        }
    }

    for _, tx := range b.Transactions {
        for _, op := range tx.SetOfOperations {
            if bch.Verify(op) == false {
                return false
            }
        }
    }

    return true
}

func (bch Blockchain) String() string {
    var output string

    for id, balance := range bch.CoinDatabase {
        output += id + ": " + fmt.Sprintf("%d", balance) + "\n"
    }
    for _, b := range bch.BlockHistory {
        output += b.String()
    }
    return output
}
