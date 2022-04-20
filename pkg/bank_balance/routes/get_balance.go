package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	pb "github.com/hellokvn/bank-api-gateway/pkg/bank_balance/pb"
)

func GetBalance(ctx *fiber.Ctx, c pb.BankBalanceQueryServiceClient) error {
	res, err := c.GetBalance(context.Background(), &pb.GetBalanceRequest{
		Id: ctx.Params("id"),
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
