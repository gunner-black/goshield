package envelope

type Envelope struct {
    EncryptedData []byte `json:"encrypted_data"`
    EncryptedKey  []byte `json:"encrypted_key"`
    Nonce         []byte `json:"nonce"`
    Algorithm     string `json:"algorithm"`
}