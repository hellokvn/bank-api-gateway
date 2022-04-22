package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	pb "github.com/hellokvn/bank-api-gateway/pkg/bank_account/pb"
)

// type AccountType string

// const (
// 	Savings AccountType = "SAVINGS"
// 	Current             = "CURRENT"
// )

type OpenAccountRequestBody struct {
	Holder         string `json:"holder"`
	Type           string `json:"type"`
	Email          string `json:"email"`
	OpeningBalance int32  `json:"openingBalance"`
}

func OpenAccount(ctx *fiber.Ctx, c pb.BankAccountCommandServiceClient) error {
	body := OpenAccountRequestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := c.OpenAccount(context.Background(), &pb.OpenAccountRequest{
		Holder:         body.Holder,
		Type:           body.Type,
		Email:          body.Email,
		OpeningBalance: body.OpeningBalance,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
