package bank_balance

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/bank-api-gateway/pkg/bank_balance/routes"
	"github.com/hellokvn/bank-api-gateway/pkg/common/config"
)

func RegisterRoutes(app *fiber.App, c config.Config) {
	svc := InitServiceClient(&c)

	r := app.Group("/bank-balance")
	r.Put("/deposit-funds/:id", svc.DepositFunds)
	r.Put("/withdraw-funds/:id", svc.WithdrawFunds)
	r.Put("/transfer-funds/:id", svc.TransferFunds)
	r.Get("/get-balance/:id", svc.GetBalance)
}

func (svc *ServiceClient) DepositFunds(ctx *fiber.Ctx) error {
	return routes.DepositFunds(ctx, svc.CommandClient)
}

func (svc *ServiceClient) WithdrawFunds(ctx *fiber.Ctx) error {
	return routes.WithdrawFunds(ctx, svc.CommandClient)
}

func (svc *ServiceClient) TransferFunds(ctx *fiber.Ctx) error {
	return routes.TransferFunds(ctx, svc.CommandClient)
}

func (svc *ServiceClient) GetBalance(ctx *fiber.Ctx) error {
	return routes.GetBalance(ctx, svc.QueryClient)
}
