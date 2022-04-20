package bank_account

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/bank-api-gateway/pkg/bank_account/routes"
	"github.com/hellokvn/bank-api-gateway/pkg/common/config"
)

func RegisterRoutes(app *fiber.App, c config.Config) {
	svc := InitServiceClient(&c)

	r := app.Group("/bank-account")
	r.Get("/find-one/:id", svc.FindAccount)
	r.Get("/find/:page", svc.FindAllAccounts)
	r.Post("/", svc.OpenAccount)
}

func (svc *ServiceClient) FindAccount(ctx *fiber.Ctx) error {
	return routes.FindAccount(ctx, svc.QueryClient)
}

func (svc *ServiceClient) FindAllAccounts(ctx *fiber.Ctx) error {
	return routes.FindAllAccounts(ctx, svc.QueryClient)
}

func (svc *ServiceClient) OpenAccount(ctx *fiber.Ctx) error {
	return routes.OpenAccount(ctx, svc.CommandClient)
}
