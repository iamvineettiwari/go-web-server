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

	router := httpServer.GetRouter()

	router.Post("/users", handlers.CreateUser)
	router.Get("/users", handlers.GetUsers)
	router.Get("/users/:id", handlers.GetUserById)
	router.Put("/users/:id", handlers.UpdateUser)
	router.Delete("/users/:id", handlers.DeleteUser)

	fmt.Println("Server starting at : http://localhost:8001")

	err := httpServer.Start()

	if err != nil {
		log.Fatal(err)
		return
	}
}
