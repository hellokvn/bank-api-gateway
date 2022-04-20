proto:
	protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.

server:
	go run cmd/main.go

proto.update:
	rm -rf pkg/common/proto && curl -L -O https://github.com/hellokvn/bank-shared-proto/archive/main.zip && unzip main.zip && rm -rf main.zip && mv bank-shared-proto-main/proto pkg/common/ && rm -rf bank-shared-proto-main