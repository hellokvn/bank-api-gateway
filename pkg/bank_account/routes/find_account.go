package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	pb "github.com/hellokvn/bank-api-gateway/pkg/bank_account/pb"
)

func FindAccount(ctx *fiber.Ctx, c pb.BankAccountQueryServiceClient) error {
	res, err := c.FindAccount(context.Background(), &pb.FindAccountRequest{
		Id: ctx.Params("id"),
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
