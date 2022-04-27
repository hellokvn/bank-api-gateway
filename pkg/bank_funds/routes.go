package bankFunds

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hellokvn/bank-api-gateway/pkg/bank_funds/routes"
	"github.com/hellokvn/bank-api-gateway/pkg/common/config"
)

func RegisterRoutes(app *fiber.App, c config.Config) {
	svc := InitServiceClient(&c)

	r := app.Group("/api/v1/bank-funds")
	r.Put("/deposit/:id", svc.DepositFunds)
	r.Put("/withdraw/:id", svc.WithdrawFunds)
	r.Put("/transfer/:id", svc.TransferFunds)
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
