package main

import (
	"fmt"
	"log"
	"makromusic-task/config"
	"makromusic-task/handlers"
	"makromusic-task/middlewares"
	"makromusic-task/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.NewConfig()

	if err := utils.InitializeDB(cfg.DBConnectionString); err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer utils.CloseDB()
	if err := utils.InitializeRedis(cfg.RedisAddress, cfg.RedisPassword, cfg.RedisDB); err != nil {
		log.Fatalf("Error initializing Redis: %v", err)
	}
	defer utils.CloseRedis()
	app := fiber.New()

	app.Use(middlewares.AuthMiddleware)

	app.Post("/register", handlers.RegisterHandler)
	app.Post("/login", handlers.LoginHandler)

	authenticated := app.Group("/api")
	authenticated.Get("/todos", handlers.GetTodoListHandler)
	authenticated.Post("/todos", handlers.AddTodoHandler)
	authenticated.Patch("/todos/:id/complete", handlers.CompleteTodoHandler)

	port := ":3000"
	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(app.Listen(port))
}
