run:
	go run cmd/main.go

build:
	go build -o chat-app cmd/main.go

docker-up:
	docker-compose up --build -d

docker-down:
	docker-compose down