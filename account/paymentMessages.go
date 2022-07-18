package account

import (
    "crypto/ed25519"
    "encoding/binary"
)

type PaymentMessages struct {
    Sender      ed25519.PublicKey
    Receiver    ed25519.PublicKey
    MoneyAmount uint64
}

func (message PaymentMessages) Encode() []byte {
    output := make([]byte, 72)
    for i := 0; i < 32; i++ {
        output[i] = message.Sender[i]
    }
    for i := 32; i < 64; i++ {
        output[i] = message.Receiver[i-32]
    }

    buf := make([]byte, 8)
    binary.PutUvarint(buf, message.MoneyAmount)

    for i := 64; i < 72; i++ {
        output[i] = buf[i-64]
    }

    return output
}

func (message *PaymentMessages) Decode(input []byte) {
    message.Sender = append(message.Sender, input[0:32]...)
    message.Receiver = append(message.Receiver, input[32:64]...)
    message.MoneyAmount, _ = binary.Uvarint(input[64:72])
}
