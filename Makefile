.PHONY: server-up
server-up:
	docker compose exec go-graphql go run server.go

.PHONY: gqlgen-init
gqlgen-init:
	cd go && go run github.com/99designs/gqlgen init
