package main

import (
    "encoding/json"
    "fmt"
    "log"

    "github.com/seu-usuario/goshield/pkg/crypto"
    "github.com/seu-usuario/goshield/pkg/provider"
)

// local test
func main() {
    fmt.Println("--- Iniciando Sistema GoShield ---")

    minhaChaveMestra := "12345678901234567890123456789012" 
    
    localVault, err := provider.NewLocal(minhaChaveMestra)
    if err != nil {
        log.Fatal(err)
    }

    dadoSecreto := []byte("Meu CPF é 123.456.789-00")

    meuEnvelope, err := crypto.Encrypt(dadoSecreto, localVault)
    if err != nil {
        log.Fatal("Erro ao criptografar:", err)
    }

    jsonOutput, _ := json.MarshalIndent(meuEnvelope, "", "  ")
    fmt.Println("O que será salvo no Banco de Dados:")
    fmt.Println(string(jsonOutput))
}
