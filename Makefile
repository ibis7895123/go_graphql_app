.PHONY: server-up
server-up:
	docker compose exec go-graphql go run server.go