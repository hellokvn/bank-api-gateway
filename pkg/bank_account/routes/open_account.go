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
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	OpeningBalance int32  `json:"openingBalance"`
	Type           string `json:"type"`
}

func OpenAccount(ctx *fiber.Ctx, c pb.BankAccountCommandServiceClient) error {
	body := OpenAccountRequestBody{}

	if err := ctx.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := c.OpenAccount(context.Background(), &pb.OpenAccountRequest{
		FirstName:      body.FirstName,
		LastName:       body.LastName,
		Type:           body.Type,
		OpeningBalance: body.OpeningBalance,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
