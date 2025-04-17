package main

import (
    "client/internal/services"
    "client/internal/transport"
    "client/internal/console"
)

func main() {
    client := transport.New("http://localhost:8080")
    analyzer := services.New(client)
    console := console.New(analyzer)
    console.Run()
}