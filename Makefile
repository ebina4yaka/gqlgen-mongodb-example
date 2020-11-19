generate:
	go run github.com/99designs/gqlgen && go run ./model_tags.go
test:
	go test ./tests
