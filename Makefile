run:
	go run cmd/main.go
build:
	docker-compose up --build
compose:
	docker-compose up -d