postgres:
	sudo docker run --name postgres15 --network auth-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

migrateup:
	migrate -path pkg/db/migration -database "postgres://root:secret@localhost:5432/auth_user?sslmode=disable" -verbose up

migratedown:
	migrate -path pkg/db/migration -database "postgres://root:secret@localhost:5432/auth_user?sslmode=disable" -verbose down

sqlc:
	sudo docker run --rm -v ${CURDIR}/pkg/db:/src -w /src kjconroy/sqlc:1.14.0 generate

server:
	go run cmd/main.go

evans:
	 evans pkg/pb/auth.proto --host localhost --port 8080