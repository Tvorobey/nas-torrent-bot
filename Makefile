.PHONY: fmt
fmt:
	go fmt ./internal/...
	goimports -w .

generate:
	go generate ./internal/...

test:
	go test -v -count 1 -race ./internal/...

lint:
	golangci-lint run

vendor:
	go mod tidy
	go mod vendor

build:
	rm torrent_bot && rm torrentbot.tar
	CGO_ENABLED=0 GOOS=linux go build -a -o ./torrent_bot ./cmd/main.go
	docker build --tag torrentbot .
	docker save -o ./torrentbot.tar torrentbot