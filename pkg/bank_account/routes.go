package bank_account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/bank-api-gateway/pkg/bank_account/routes"
	"github.com/hellokvn/bank-api-gateway/pkg/common/config"
)

func RegisterRoutes(app *fiber.App, c config.Config) {
	svc := &ServiceClient{
		Client: InitServiceClient(&c),
	}

	r := app.Group("/bank-account")
	r.Get("/id/:id", svc.FindAccount)
}

func (svc *ServiceClient) FindAccount(ctx *fiber.Ctx) error {
	return routes.FindAccount(ctx, svc.Client)
}
