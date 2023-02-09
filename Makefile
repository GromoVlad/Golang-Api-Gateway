.PHONY: run
run:
	swag init -g cmd/main.go && docker-compose up && goose-i && go get google.golang.org/grpc
migration-status:
	docker exec gin_tonic $(MAKE) migration-status -C migration
migration-up:
	docker exec gin_tonic $(MAKE) migration-up -C migration
migration-down:
	docker exec gin_tonic $(MAKE) migration-down -C migration
migration-create:
	docker exec gin_tonic $(MAKE) migration-create -C migration
documentation-create:
	swag init -g cmd/main.go
test:
	docker exec gin_tonic go test ./tests/...
goose-i:
	curl -fsSL \
	    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
	    sh
refresh-proto:
	protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  		internal/controllers/apiGateway/book/listBook/gRPC/listBook.proto && \
		protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		internal/controllers/apiGateway/book/findBook/gRPC/findBook.proto && \
		protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		  internal/controllers/apiGateway/book/createBook/gRPC/createBook.proto && \
		protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		  internal/controllers/apiGateway/book/updateBook/gRPC/updateBook.proto && \
		protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		   internal/controllers/apiGateway/book/deleteBook/gRPC/deleteBook.proto
