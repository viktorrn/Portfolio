run:
	@templ generate
	@go build src/main.go
	@./main