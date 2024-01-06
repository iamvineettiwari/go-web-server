package main

import (
	"fmt"
	"log"

	"github.com/iamvineettiwari/go-web-server/server"
)

func main() {
	httpServer := server.NewHttpServer("127.0.0.1:8001")

	fmt.Println("Server starting at : http://localhost:8001")

	err := httpServer.Start()

	if err != nil {
		log.Fatal(err)
		return
	}
}
