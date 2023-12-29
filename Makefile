dummy:
	@go build -o bin/dummy twitter-api/main.go
	@./bin/dummy

fetcher:
	@go build -o bin/fetcher fetcher/main.go
	@./bin/fetcher

.PHONY: dumy fetcher
