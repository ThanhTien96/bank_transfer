postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb: 
	sudo docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	sudo docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate


gotest:
	go test -v -cover ./...

.PONY: postgres createdb dropdb migrateup migratedown