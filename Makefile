.PHONY: run
run:
	swag init -g cmd/main.go && docker-compose up && goose-i
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
	go test ./tests/...
goose-i:
	curl -fsSL \
	    https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
	    sh