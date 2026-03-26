run:
	go run cmd/api/main.go

migration:
	@read -p "migration file name: " file;\
	go run cmd/migration/main.go -file=$$file

seed-hospital:
	go run cmd/seed/main.go

compose-up:
	cp .env.dev .env
	docker compose up -d

compose-down:
	docker-compose down