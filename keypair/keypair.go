package keypair

import (
    "blokchain/signature"
    "crypto/ed25519"
)

type KeyPair struct {
    privateKey ed25519.PrivateKey
    PublicKey  ed25519.PublicKey
}

func (k *KeyPair) GenerateKeyPair() error {
    var err error
    k.PublicKey, k.privateKey, err = ed25519.GenerateKey(nil)
    return err
}

func (k KeyPair) SignData(message []byte) signature.SignedMessage {
    var output signature.SignedMessage
    output.Signature = ed25519.Sign(k.privateKey, message)
    output.Data = message
    return output
}
