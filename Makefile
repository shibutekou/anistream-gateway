.PHONY: start dr dc redc protoc gw

start:
	go build -o bin/anistream-gateway cmd/main.go
	./bin/anistream-gateway

dr:
	docker build --tag anistream-gateway .
	docker run --network=host anistream-gateway

dc:
	docker compose down
	docker compose up

redc:
	docker compose down
	docker compose up --build

protoc:
	protoc --go_out=./internal/service/content/pb --go_opt=paths=source_relative \
    --go-grpc_out=./internal/service/content/pb --go-grpc_opt=paths=source_relative \
    --proto_path=./proto proto/content.proto

gw:
	protoc --proto_path=proto/ --grpc-gateway_out ./internal/service/content/pb \
         --grpc-gateway_opt logtostderr=true \
         --grpc-gateway_opt paths=source_relative \
         --grpc-gateway_opt generate_unbound_methods=true \
          ./proto/content.proto
