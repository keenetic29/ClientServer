package main

import (
	"server/internal/routes"
)

func main() {
	r := routes.SetupRouter()
	r.Run(":8080")
}