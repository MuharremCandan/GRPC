
dev:
	docker compose up -d && air

proto:
	rm -f pb/*.go  
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

evans:
	evans --host 127.0.0.1 --port 9090 --proto pkg/proto/user.proto repl

.PHONY: dev proto testgrpc