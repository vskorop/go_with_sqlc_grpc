postgres:
	docker run --name grpc-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pedro1 -d postgres:12-alpine

createdb: 
	docker exec -it grpc-postgres createdb --username=root --owner=root simple_bank

dropdb: 
	docker exec -it grpc-postgres dropdb simple_bank

migrateup: 
	migrate -path db/migration -database "postgresql://root:pedro1@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://root:pedro1@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc: 
	sqlc generate

.PHONY: posrgres createdb dropdb migrateup migratedown sqlc