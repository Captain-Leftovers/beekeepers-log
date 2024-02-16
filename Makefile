DEV:
	@templ generate
	@go build -o ./tmp/main ./cmd/main.go
	@./tmp/main