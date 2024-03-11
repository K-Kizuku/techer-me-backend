.PHONY: lint lint-fix run logs migrate migrate-down gen-migrate

lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run ./...  --fix

run:
	docker compose up --build -d

logs:
	docker compose logs -f

migrate:
	migrate -path db/migrations -database "mysql://user:password@tcp(localhost:3306)/db?multiStatements=true" up

migrate-down:
	migrate -path db/migrations -database "mysql://user:password@tcp(localhost:3306)/db?multiStatements=true" down

gen-migrate:
	migrate create -ext sql -dir db/migrations -seq $(name)