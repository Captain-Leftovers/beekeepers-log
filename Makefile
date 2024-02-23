

templ:
	templ generate -watch -proxy=http://localhost:3000

tailwind:
	tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch

DEV:
	@templ generate
	@go build -o ./tmp/main ./cmd/main.go
	@./tmp/main



