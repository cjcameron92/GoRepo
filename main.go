package main

import (
	"fmt"
	"gorepo/server/server"
)

func main() {
	fmt.Println("Listening on port 8080")
	server.StartServer(":8080", "./repo")
}
