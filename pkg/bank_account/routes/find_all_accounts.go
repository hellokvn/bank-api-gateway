package routes

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	pb "github.com/hellokvn/bank-api-gateway/pkg/bank_account/pb"
)

func FindAllAccounts(ctx *fiber.Ctx, c pb.BankAccountQueryServiceClient) error {
	page, err := strconv.ParseInt(ctx.Params("page"), 10, 32)

	if err != nil || page < 1 {
		page = 1
	}

	res, err := c.FindAllAccounts(context.Background(), &pb.FindAllAccountsRequest{
		Page: int32(page),
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())
	}

	return ctx.Status(int(res.Status)).JSON(&res)
}
