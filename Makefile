run:
	go run cmd/redump/main.go

set-mock-redmine:
	cd docker/mock_redmine && docker-compose up -d
