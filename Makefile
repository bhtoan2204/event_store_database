up:
	docker compose up -d --build

down:
	docker compose down

run-user:
	cd user && go run cmd/main.go