package bankFunds

import (
	"fmt"

	pb "github.com/hellokvn/bank-api-gateway/pkg/bank_funds/pb"
	"github.com/hellokvn/bank-api-gateway/pkg/common/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	CommandClient pb.BankFundsCommandServiceClient
	QueryClient   pb.BankFundsQueryServiceClient
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
		QueryClient:   pb.NewBankFundsQueryServiceClient(queryConnection),
		CommandClient: pb.NewBankFundsCommandServiceClient(cmdConnection),
	}
}
