package main

import (
	"fmt"
	"log"

	"github.com/iamvineettiwari/go-web-server/handlers"
	"github.com/iamvineettiwari/go-web-server/server"
)

func main() {
	httpServer := server.NewHttpServer("127.0.0.1:8001")

	// httpServer.SetStaticPath(filepath.Join(server.BASE_PATH, "cmd/bin/www"))

	httpServer.Get("/users", handlers.GetUser)
	httpServer.Post("/users", handlers.CreateUser)
	httpServer.Put("/users", handlers.UpdateUser)
	httpServer.Delete("/users", handlers.DeleteUser)

	fmt.Println("Server starting at : http://localhost:8001")

	err := httpServer.Start()

	if err != nil {
		log.Fatal(err)
		return
	}
}
