package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	pb "github.com/hellokvn/bank-api-gateway/pkg/bank_funds/pb"
)

type TransferFundsRequestBody struct {
	Amount int32  `json:"amount"`
	ToId   string `json:"toId"`
}

func TransferFunds(ctx *fiber.Ctx, c pb.BankFundsCommandServiceClient) error {
	body := TransferFundsRequestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := c.TransferFunds(context.Background(), &pb.TransferFundsRequest{
		FromId: ctx.Params("id"),
		ToId:   body.ToId,
		Amount: body.Amount,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
