
dev:
	docker compose up -d && air

proto:
	rm -f pkg/pb/*.go  
	protoc --proto_path=pkg/proto --go_out=pkg/pb --go_opt=paths=source_relative \
    --go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pkg/pb --grpc-gateway_opt=paths=source_relative \
    pkg/proto/*.proto

evans:
	evans --host 127.0.0.1 --port 9090 --proto pkg/proto/user.proto repl

.PHONY: dev proto evans