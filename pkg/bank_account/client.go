package bankAccount

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
	queryConnection, err := grpc.Dial(c.BankAccountQSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	cmdConnection, err := grpc.Dial(c.BankAccountCSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return &ServiceClient{
		QueryClient:   pb.NewBankAccountQueryServiceClient(queryConnection),
		CommandClient: pb.NewBankAccountCommandServiceClient(cmdConnection),
	}
}
