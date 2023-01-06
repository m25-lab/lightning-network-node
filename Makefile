proto:
	rm -f rpc/pb/*.go
	protoc --proto_path=rpc/proto --go_out=rpc/pb --go_opt=paths=source_relative \
    --go-grpc_out=rpc/pb --go-grpc_opt=paths=source_relative \
    rpc/proto/*.proto

.PHONY: proto
