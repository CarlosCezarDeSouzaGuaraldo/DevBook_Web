package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// HOST is the application host
	HOST = ""

	// PORT is the port from API url
	PORT = 0

	// URL_API used to represent the API url
	URL_API = ""

	// HASH_KEY is used to authenticate the cookie
	HASH_KEY []byte

	// BLOCK_KEY is used to encrypte the cookie data
	BLOCK_KEY []byte
)

// Load initialize the enviroment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}

	HOST = os.Getenv("HOST")
	URL_API = os.Getenv("URL_API")
	HASH_KEY = []byte(os.Getenv("HASH_KEY"))
	BLOCK_KEY = []byte(os.Getenv("BLOCK_KEY"))
}
