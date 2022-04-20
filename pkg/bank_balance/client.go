package bank_balance

import (
	"fmt"

	pb "github.com/hellokvn/bank-api-gateway/pkg/bank_balance/pb"
	"github.com/hellokvn/bank-api-gateway/pkg/common/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	CommandClient pb.BankBalanceCommandServiceClient
	QueryClient   pb.BankBalanceQueryServiceClient
}

func InitServiceClient(c *config.Config) *ServiceClient {
	cc, err := grpc.Dial(c.BankAccountQSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &ServiceClient{
		QueryClient:   pb.NewBankBalanceQueryServiceClient(cc),
		CommandClient: pb.NewBankBalanceCommandServiceClient(cc),
	}
}
