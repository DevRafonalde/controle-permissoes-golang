package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
)

func GeraChaveSecreta() error {
	chave := make([]byte, 32)
	_, err := rand.Read(chave)
	if err != nil {
		return fmt.Errorf("Falha ao gerar a chave secreta: %w", err)
	}

	chaveCodificada := base64.StdEncoding.EncodeToString(chave)

	// Caminho da pasta e do arquivo
	folderPath := "./jwt"
	fileName := "/jwt_secret_key.txt"
	filePath := folderPath + fileName

	// Verifica se a pasta existe e, se não, cria
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			return err
		}
	}

	// Write the key to a file
	err = os.WriteFile(filePath, []byte(chaveCodificada), 0644)
	if err != nil {
		return fmt.Errorf("Falha ao escrever a chave secreta em um arquivo: %w", err)
	}

	return nil
}
