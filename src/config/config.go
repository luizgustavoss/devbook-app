package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	// ApiUrl URL of the back-end API
	ApiUrl string
	// Port API port number
	Port     int
	// HashKey used to authenticate cookie
	HashKey  []byte
	// BlockKey used to encrypt the cookie
	BlockKey []byte
)

// Load inits environment variables
func Load() {
	var error error

	if error = godotenv.Load(); error != nil {
		log.Fatal(error)
	}

	Port, error = strconv.Atoi(os.Getenv("API_PORT"))
	if error != nil {
		Port = 9000 // assumes default port number
	}

	ApiUrl = os.Getenv("API_URL")
	if error != nil {
		log.Fatal(error)
	}

}
