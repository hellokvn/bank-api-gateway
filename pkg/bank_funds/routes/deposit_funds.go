package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	pb "github.com/hellokvn/bank-api-gateway/pkg/bank_funds/pb"
)

type DepositFundsRequestBody struct {
	Amount int32 `json:"amount"`
}

func DepositFunds(ctx *fiber.Ctx, c pb.BankFundsCommandServiceClient) error {
	body := DepositFundsRequestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := c.DepositFunds(context.Background(), &pb.DepositFundsRequest{
		Id:     ctx.Params("id"),
		Amount: body.Amount,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
