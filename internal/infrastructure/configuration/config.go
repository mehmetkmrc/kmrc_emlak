package configuration

import (
	
	"os"
)

type Config struct {
    ServerPort string
}

func LoadConfig() Config {
    port := os.Getenv("SERVER_PORT")
    if port == "" {
        port = ":3000" // Varsayılan port
    }
    return Config{
        ServerPort: port,
    }
}
