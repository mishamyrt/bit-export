package main

import (
	"bitexporter/internal/api"
	"bitexporter/internal/crypto"
	"bitexporter/internal/domain"
	"bitexporter/internal/export"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func getState(apiUrl string, id string, secret string) (*domain.Sync, domain.Auth) {
	api, err := api.New(apiUrl, id, secret)
	if err != nil {
		log.Fatalf("Authorization error: %v", err)
	}
	sync, err := api.Sync()
	if err != nil {
		log.Fatalf("Synchronization error: %v", err)
	}
	return &sync, api.Auth
}

func decryptState(sync *domain.Sync, key string, password string, params crypto.KDFParams) {
	userKey, err := crypto.CalculateUserKey(password, sync.Profile.Email, params)
	masterKey, macKey, err := crypto.DecryptMasterKey([]byte(key), userKey)
	if err != nil {
		log.Fatalf("Master key decryption error: %v", err)
	}
	var coder crypto.Coder
	coder.SetKeys(masterKey, macKey)
	coder.DecryptSync(sync)
}

func exportState(sync *domain.Sync) {
	file := export.FromDomain(sync)
	content, err := json.Marshal(&file)
	if err != nil {
		log.Fatalf("JSON marshall error: %v", err)
	}
	jsonContent := string(content)
	fmt.Println(jsonContent)
	err = ioutil.WriteFile("bw-export.json", []byte(jsonContent), 0644)
	if err != nil {
		log.Fatalf("File writing error: %v", err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientSecret := os.Getenv("BW_CLIENT_SECRET")
	clientId := os.Getenv("BW_CLIENT_ID")
	password := os.Getenv("BW_PASSWORD")
	apiUrl := os.Getenv("BW_API_URL")

	sync, auth := getState(apiUrl, clientId, clientSecret)
	decryptState(sync, auth.Key, password, crypto.KDFParams{
		Type:        crypto.KDF(auth.Kdf),
		Memory:      auth.KdfMemory,
		Iterations:  auth.KdfIterations,
		Parallelism: auth.KdfParallelism,
	})

	exportState(sync)
}
