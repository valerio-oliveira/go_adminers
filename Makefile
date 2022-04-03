killconn:
	psql -d postgres -c "SELECT pg_terminate_backend( pid ) FROM pg_stat_activity WHERE pid <> pg_backend_pid( ) AND datname in ('adminers');"

dropdb:
	psql -d postgres -c "DROP DATABASE adminers;"

createdb:
	psql -d postgres -c "CREATE DATABASE adminers;"

migrateup:
	migrate -path db/migration/ -database "postgresql://valerio:admin@123@localhost:5432/adminers?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://valerio:admin@123@localhost:5432/adminers?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

all: killconn dropdb createdb migrateup sqlc test

.PHONY: dropdb createdb migrateup migratedown sqlc test

