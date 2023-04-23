package main

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	cache "github.com/smerschjohann/act-cache/artifactcache"
)

func main() {
	err := godotenv.Load(".actcacherc")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		log.Fatal("Error loading .actcacherc file")
	}

	cacheDir := os.Getenv("ACT_CACHE_DIRECTORY")
	if cacheDir == "" {
		cacheDir = ".cache"
	}
	cacheIp := os.Getenv("ACT_CACHE_IP") // StartHandler gets IP automatically if not set

	cachePort := 8111
	cachePortString := os.Getenv("ACT_CACHE_PORT")
	if cachePortString != "" {
		parsedPort, err := strconv.Atoi(cachePortString)
		if err != nil {
			log.Fatalf("could not parse port: %v", err)
		}
		cachePort = parsedPort
	}

	handler, err := cache.StartHandler(cacheDir, cacheIp, uint16(cachePort))

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on: %s", handler.ExternalURL())
	handler.WaitForever()
}
