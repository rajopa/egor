.PHONY: migrate-up migrate-down

migrate-up:
	migrate -path ./schema -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" up

migrate-down:
	migrate -path ./schema -database "postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable" down
deploy:
	kubectl apply -f k8s/