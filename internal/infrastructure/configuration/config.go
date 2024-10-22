package configuration

import (
    "os"
)

type Config struct {
    ServerPort string
}

func LoadConfig() Config {
    // Portu çevre değişkenlerinden al, varsayılan olarak ":3000" kullan
    port := os.Getenv("SERVER_PORT")
    if port == "" {
        port = ":3000"
    }

    return Config{
        ServerPort: port,
    }
}
