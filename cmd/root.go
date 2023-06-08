// Package cmd contains descriptions and handlers for vpn-dns CLI.
package cmd

import (
	"bit-exporter/internal/api"
	"bit-exporter/internal/crypto"
	"bit-exporter/internal/domain"
	"bit-exporter/internal/export"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

// AppName represents app name.
const AppName = "bit-exporter"

// Version represents current app version.
var Version = "development"

var clientSecret string
var clientId string
var password string
var apiUrl string

func getEnvAssert(key string, target *string) {
	*target = os.Getenv(key)
	if len(*target) == 0 {
		log.Fatalf("$%v variable is not set", key)
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	getEnvAssert("BW_CLIENT_SECRET", &clientSecret)
	getEnvAssert("BW_CLIENT_ID", &clientId)
	getEnvAssert("BW_PASSWORD", &password)
	getEnvAssert("BW_API_URL", &apiUrl)
}

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

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:     AppName,
	Version: Version,
	Short:   "The app that exports records from a Bitwarden-compatible server",
	Run: func(cmd *cobra.Command, args []string) {
		sync, auth := getState(apiUrl, clientId, clientSecret)
		decryptState(sync, auth.Key, password, crypto.KDFParams{
			Type:        crypto.KDF(auth.Kdf),
			Memory:      auth.KdfMemory,
			Iterations:  auth.KdfIterations,
			Parallelism: auth.KdfParallelism,
		})

		exportState(sync)
	},
}

// Execute is the main CLI entrypoint.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
