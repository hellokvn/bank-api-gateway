package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	pb "github.com/hellokvn/bank-api-gateway/pkg/bank_funds/pb"
)

type WithdrawFundsRequestBody struct {
	Amount int32 `json:"amount"`
}

func WithdrawFunds(ctx *fiber.Ctx, c pb.BankFundsCommandServiceClient) error {
	body := WithdrawFundsRequestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := c.WithdrawFunds(context.Background(), &pb.WithdrawFundsRequest{
		Id:     ctx.Params("id"),
		Amount: body.Amount,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
