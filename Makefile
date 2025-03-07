.PHONY: build run migrate

build:
	docker-compose build

run:
	docker-compose up

migrate:
	docker-compose exec db psql -U postgres -d library -f /app/migrations/001_create_books_table.sql
