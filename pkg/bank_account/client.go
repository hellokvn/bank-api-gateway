package bank_account

import (
	"fmt"

	pb "github.com/hellokvn/bank-api-gateway/pkg/bank_account/pb"
	"github.com/hellokvn/bank-api-gateway/pkg/common/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.BankAccountQueryServiceClient
}

func InitServiceClient(c *config.Config) pb.BankAccountQueryServiceClient {
	cc, err := grpc.Dial(c.BankAccountQSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewBankAccountQueryServiceClient(cc)
}
