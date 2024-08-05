.PHONY: tidy
tidy:
	go fmt ./internal/...

generate:
	go generate ./internal/...

test:
	go test -v -count 1 -race ./internal/...

build:
	rm torrent_bot && rm torrentbot.tar
	CGO_ENABLED=0 GOOS=linux go build -a -o ./torrent_bot ./cmd/main.go
	docker build --tag torrentbot .
	docker save -o ./torrentbot.tar torrentbot