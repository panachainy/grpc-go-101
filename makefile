path = pingpong-all.proto

gen.proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    $(path)

unary.server:
	go run ./cmd/unary/server/
unary.client:
	go run ./cmd/unary/client/

