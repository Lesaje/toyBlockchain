package signature

import (
    "crypto/ed25519"
)

type SignedMessage struct {
    Data      []byte
    Signature []byte
}

func (message SignedMessage) Verify(publicKey ed25519.PublicKey) bool {
    return ed25519.Verify(publicKey, message.Data, message.Signature)
}
