package bank_account

import (
	"fmt"

	pb "github.com/hellokvn/bank-api-gateway/pkg/bank_account/pb"
	"github.com/hellokvn/bank-api-gateway/pkg/common/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	CommandClient pb.BankAccountCommandServiceClient
	QueryClient   pb.BankAccountQueryServiceClient
}

func InitServiceClient(c *config.Config) *ServiceClient {
	cc, err := grpc.Dial(c.BankAccountQSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &ServiceClient{
		QueryClient:   pb.NewBankAccountQueryServiceClient(cc),
		CommandClient: pb.NewBankAccountCommandServiceClient(cc),
	}
}
