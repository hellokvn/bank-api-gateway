package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	bankAccount "github.com/hellokvn/bank-api-gateway/pkg/bank_account"
	bankFunds "github.com/hellokvn/bank-api-gateway/pkg/bank_funds"
	"github.com/hellokvn/bank-api-gateway/pkg/common/config"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()

	app.Use(cors.New())

	bankAccount.RegisterRoutes(app, c)
	bankFunds.RegisterRoutes(app, c)

	err = app.Listen(fmt.Sprintf(":%s", c.Port))

	if err != nil {
		log.Fatalln("Failed to listen", err)
	}
}
