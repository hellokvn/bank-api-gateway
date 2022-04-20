server:
	go run cmd/main.go

proto.update:
	rm -rf pkg/common/proto && curl -L -O https://github.com/hellokvn/bank-shared-proto/archive/main.zip && unzip main.zip && rm -rf main.zip && mv bank-shared-proto-main/proto pkg/common/ && rm -rf bank-shared-proto-main

proto.bank-account:
	protoc pkg/common/proto/bank-account-*.proto --go_out=plugins=grpc:./pkg/bank_account/pb/

proto.bank-balance:
	protoc pkg/common/proto/bank-balance-*.proto --go_out=plugins=grpc:./pkg/bank_balance/pb/

proto:
	make proto.update && make proto.bank-account && make proto.bank-balance
