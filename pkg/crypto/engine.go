package crypto

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "io"
    "github.com/gunner-black/goshield/pkg/envelope"
    "github.com/gunner-black/goshield/pkg/provider"
)

func Encrypt(data []byte, master provider.MasterKeyProvider) (*envelope.Envelope, error) {
    dek := make([]byte, 32)
    if _, err := io.ReadFull(rand.Reader, dek); err != nil {
        return nil, err 
    }

    block, err := aes.NewCipher(dek)
    if err != nil {
        return nil, err
    }
    
    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }

    nonce := make([]byte, aesGCM.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }

    encryptedData := aesGCM.Seal(nil, nonce, data, nil)

    encryptedDEK, err := master.EncryptDEK(dek)
    if err != nil {
        return nil, err
    }

    return &envelope.Envelope{
        EncryptedData: encryptedData,
        EncryptedKey:  encryptedDEK,
        Nonce:         nonce,
        Algorithm:     "AES-256-GCM",
    }, nil
}