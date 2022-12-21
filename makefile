run: stop up

mod:
	# This make rule requires Go 1.19+
	GO119MODULE=on go mod tidy

up:
	docker-compose -f docker-compose.yml up -d --build

stop:
	docker-compose -f docker-compose.yml stop

down:
	docker-compose -f docker-compose.yml down

migrate-up:
	migrate --path src/migrations -database "postgresql://postgres:Password!123#@localhost:5432?sslmode=disable" -verbose up

migrate-down:
	migrate --path src/migrations -database "postgresql://postgres:Password!123#@localhost:5432?sslmode=disable" -verbose down