proto:
	rm -f node/rpc/pb/*.go
	protoc --proto_path=node/rpc/proto --go_out=node/rpc/pb --go_opt=paths=source_relative \
    --go-grpc_out=node/rpc/pb --go-grpc_opt=paths=source_relative \
    node/rpc/proto/*.proto

.PHONY: proto
