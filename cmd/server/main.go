package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/versenilvis/sendio/internal/database"
	"github.com/versenilvis/sendio/internal/utils"
)

func main() {
	psqlDB, err := database.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer utils.Close(psqlDB)

	redisDB, err := database.NewRedisConnection()
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}
	defer utils.Close(redisDB)

	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	log.Println("Server starting on localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
