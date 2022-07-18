package account

import (
    "blokchain/signature"
    "crypto/ed25519"
    "encoding/binary"
    "errors"
    "fmt"
)

type Operation struct {
    Sender        ed25519.PublicKey
    Receiver      ed25519.PublicKey
    Amount        uint64
    SignedMessage signature.SignedMessage
}

func CreateOp(sender ed25519.PublicKey, message signature.SignedMessage) (Operation, error) {

    var output Operation
    var paymentData PaymentMessages
    if message.Verify(sender) != true {
        return output, errors.New("invalid signature")
    }

    bytePaymentData := message.Data
    paymentData.Decode(bytePaymentData)

    output.Amount = paymentData.MoneyAmount
    output.Sender = paymentData.Sender
    output.Receiver = paymentData.Receiver
    output.SignedMessage = message
    return output, nil
}

func (op Operation) ToByte() []byte {
    var output []byte
    output = append(output, op.Sender...)
    output = append(output, op.Receiver...)
    buf := make([]byte, 8)
    binary.PutUvarint(buf, op.Amount)
    output = append(output, buf...)
    output = append(output, op.SignedMessage.Data...)
    output = append(output, op.SignedMessage.Signature...)
    return output
}

func (bch Blockchain) Verify(op Operation) bool {

    if op.SignedMessage.Verify(op.Sender) != true {
        return false
    }

    sender := fmt.Sprintf("%x", op.Sender)
    if op.Amount > bch.CoinDatabase[sender] {
        return false
    }

    receiver := fmt.Sprintf("%x", op.Receiver)
    if _, ok := bch.CoinDatabase[receiver]; ok == false {
        return false
    }

    return true
}

func (op Operation) String() string {
    return fmt.Sprintf("%x -> %x: %d", op.Sender, op.Receiver, op.Amount)
}
