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
	queryConnection, err := grpc.Dial(c.BankFundsQSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	cmdConnection, err := grpc.Dial(c.BankFundsCSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &ServiceClient{
		QueryClient:   pb.NewBankBalanceQueryServiceClient(queryConnection),
		CommandClient: pb.NewBankBalanceCommandServiceClient(cmdConnection),
	}
}
