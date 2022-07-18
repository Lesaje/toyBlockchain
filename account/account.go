package account

import (
    "blokchain/keypair"
    "crypto/ed25519"
    "fmt"
)

type Account struct {
    Identifier ed25519.PublicKey
    wallet     keypair.KeyPair
    Balance    uint64
}

func (bch *Blockchain) NewAccount(pair keypair.KeyPair) Account {
    var output Account
    output.Identifier = pair.PublicKey
    output.wallet = pair
    bch.CoinDatabase[fmt.Sprintf("%x", output.Identifier)] = 0
    return output
}

func (acc Account) CreateOp(receiver ed25519.PublicKey, moneyAmount uint64) (Operation, error) {

    var message PaymentMessages
    message.Sender = acc.Identifier
    message.Receiver = receiver
    message.MoneyAmount = moneyAmount

    signedMessage := acc.wallet.SignData(message.Encode())
    return CreateOp(acc.Identifier, signedMessage)
}
