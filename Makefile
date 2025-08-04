postgres:
	docker run --name postgres14 \
	-e POSTGRES_USER=root \
	-e POSTGRES_PASSWORD=secret \
	-p 5432:5432 \
	-d postgres:14

createdb:
	 docker exec -it postgres14 createdb --username=root --owner=root bankapp

dropdb:
	 docker exec -it postgres14 dropdb bankapp

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankapp?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankapp?sslmode=disable" -verbose up	1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankapp?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bankapp?sslmode=disable" -verbose down	1

sqlc:
	sqlc	generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ShacharES/BankApp/db/sqlc Store

.PHONY: createdb dropdb postgres14 migrateup migratedown migrateup1 migratedown1 sqlc test server mock