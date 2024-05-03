run: build
	@./tmp/main

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
# because im using  .exe tailwind file to avoid node_modules
	@npm install -D tailwindcss

css:
	@./tailwindcss -i view/css/app.css -o public/styles.css --watch

templ:
	@templ generate --watch --proxy='http://localhost:3000

build:
	@templ generate view
	@go build -tags dev -o tmp/main main.go 


# im developing on windows so i have to use .exe please adjust to your OS
# im running from tasks.json
