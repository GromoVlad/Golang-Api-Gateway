.PHONY: run
migration-status:
	goose postgres 'postgres://postgres:@localhost:1111/gin_tonic_two' status
migration-up:
	goose postgres 'postgres://postgres:@localhost:5432/gin_tonic_two' up
migration-down:
	goose postgres 'postgres://postgres:@localhost:1111/gin_tonic_two' down
migration-create:
	goose postgres 'postgres://postgres:@localhost:1111/gin_tonic_two' create $(name) sql