# go-web-server
Building own basic web server that support a small subset of HTTP/1.1

### Implemented Functionalities

- Static file serving

  ```
  httpServer.SetStaticPath(filepath.Join(server.BASE_PATH, "cmd/bin/www"))
  ```
  
- Dynamic routing with parameters
  
  ```
    router.Post("/users", handlers.CreateUser)
	router.Get("/users", handlers.GetUsers)
	router.Get("/users/:id", handlers.GetUserById)
	router.Put("/users/:id", handlers.UpdateUser)
	router.Delete("/users/:id", handlers.DeleteUser)
  ```
  
- Concurrent connections serve

- For more details, visit [here](https://codingchallenges.fyi/challenges/challenge-webserver/)
