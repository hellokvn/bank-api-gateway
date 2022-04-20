package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hellokvn/bank-api-gateway/pkg/bank_account"
	"github.com/hellokvn/bank-api-gateway/pkg/common/config"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()

	app.Use(cors.New())

	bank_account.RegisterRoutes(app, c)

	app.Listen(c.Port)
}
