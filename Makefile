ARGS = $1

run:
	go run cmd/redump/main.go "${ARGS}"

set-mock-redmine:
	cd docker/mock_redmine && docker-compose up -d

test:
	go test -v -cover ./...

bench:
	go test -bench ./... -benchmem -cpuprofile=*.prof

build:
	go build cmd/redump/main.go
