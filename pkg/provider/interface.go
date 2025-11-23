package provider

type MasterKeyProvider interface {
    EncryptDEK(dek []byte) ([]byte, error)
    DecryptDEK(encryptedDek []byte) ([]byte, error)
}