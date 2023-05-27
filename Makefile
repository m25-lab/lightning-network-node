proto:
	rm -f rpc/pb/*.go
	protoc --proto_path=rpc/proto --go_out=rpc/pb --go_opt=paths=source_relative \
    --go-grpc_out=rpc/pb --go-grpc_opt=paths=source_relative \
    rpc/proto/*.proto

routing:
	rm -f rpc/pb/routing*.go
	protoc --proto_path=rpc/proto --go_out=rpc/pb --go_opt=paths=source_relative \
	--go-grpc_out=rpc/pb --go-grpc_opt=paths=source_relative \
	rpc/proto/routing.proto

.PHONY: proto
