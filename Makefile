build:
	docker build -t gym-postgres:dev .

start:
	docker run --name gym-postgres-container -d -p 5432:5432 gym-postgres:dev

stop:
	docker stop gym-postgres-container
	docker container rm gym-postgres-container

migrate: start
	docker run -v $(shell pwd)/internal/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgresql://boyar:go-boyars@localhost:5432/gym?sslmode=disable up