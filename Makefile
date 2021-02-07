ARGS = $1

run:
	go run cmd/redump/main.go "${ARGS}"

set-mock-redmine:
	cd docker/mock_redmine && docker-compose up -d

stop-mock-redmine:
	cd docker/mock_redmine && docker-compose down

test:
	go test -v -cover ./...

bench:
	cd pkg/redmine && \
	go test -bench . -benchmem -cpuprofile cpu.prof -memprofile mem.pro

doc:
	cd pkg && \
	godoc -http=:8080

build:
	go build -o redump cmd/redump/main.go
