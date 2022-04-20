package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	pb "github.com/hellokvn/bank-api-gateway/pkg/bank_account/pb"
)

func CloseAccount(ctx *fiber.Ctx, c pb.BankAccountCommandServiceClient) error {
	res, err := c.CloseAccount(context.Background(), &pb.CloseAccountRequest{
		Id: ctx.Params("id"),
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
