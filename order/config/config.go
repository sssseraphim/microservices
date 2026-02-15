package config

import (
	"log"
	"os"
	"strconv"
)

func getEnviromentValue(key string) string {
	if os.Getenv(key) == "" {
		log.Fatalf("%s env variable is missing.", key)
	}
	return os.Getenv(key)
}

func GetEnv() string {
	return getEnviromentValue("ENV")
}

func GetDataSourceURL() string {
	return getEnviromentValue("DATA_SOURCE_URL")
}

func GetApplicationPort() int {
	portStr := getEnviromentValue("APPLICATION_PORT")
	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Fatalf("port: %v is invalid", port)
	}

	return port
}
