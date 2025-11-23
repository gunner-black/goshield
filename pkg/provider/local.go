package provider

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "errors"
    "io"
)

type LocalProvider struct {
    masterKey []byte
}

func NewLocal(keyString string) (*LocalProvider, error) {
    if len(keyString) != 32 {
        return nil, errors.New("a chave mestra local deve ter exatamente 32 caracteres")
    }
    return &LocalProvider{masterKey: []byte(keyString)}, nil
}

func (l *LocalProvider) EncryptDEK(dek []byte) ([]byte, error) {
    block, _ := aes.NewCipher(l.masterKey)
    gcm, _ := cipher.NewGCM(block)
    nonce := make([]byte, gcm.NonceSize())
    io.ReadFull(rand.Reader, nonce)
    return gcm.Seal(nonce, nonce, dek, nil), nil 
}

func (l *LocalProvider) DecryptDEK(encryptedDek []byte) ([]byte, error) {
    return nil, nil
}