package account

import (
    "blokchain/keypair"
    "crypto/ed25519"
)

type Depository struct {
    Identifier      ed25519.PublicKey
    wallet          keypair.KeyPair
    Balance         uint64
    DepositDatabase map[[32]byte]uint64
}

func NewDepository(pair keypair.KeyPair, faucetCoins uint64) Depository {
    var output Depository
    output.Identifier = pair.PublicKey
    output.wallet = pair
    output.Balance = faucetCoins
    output.DepositDatabase = make(map[[32]byte]uint64)
    return output
}

func (d Depository) SendInterests(rec ed25519.PublicKey, amount uint64) (Operation, error) {

    var message PaymentMessages
    message.Sender = d.Identifier
    message.Receiver = rec
    message.MoneyAmount = amount
    signedMessage := d.wallet.SignData(message.Encode())

    op, err := CreateOp(d.Identifier, signedMessage)
    return op, err
}

/*
not real function, eligible only with Proof-of-Work
similar functions should be for Account struct

This function should be called by each account/deposit when verifying a new block,
and if there were transactions in the block that indicated this account/deposit as a recipient,
then this function accrue the corresponding number of coins

func (d Depository) ReceivePayment(b Block) {
    this.DepositoryDatabase[b.Hash] = b.Amount
}
*/
