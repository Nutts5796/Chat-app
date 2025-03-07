.PHONY: build run migrate

build:
	docker-compose build

run:
	docker-compose up

down:
	docker-compose down

migrate:
	docker-compose exec db psql -U postgres -d library -f /app/migrations/001_create_books_table.sql
