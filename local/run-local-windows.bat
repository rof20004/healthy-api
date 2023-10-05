set DATABASE_HOST=127.0.0.1
set DATABASE_PORT=5432
set DATABASE_USER=postgres
set DATABASE_PASS=postgres
set DATABASE_NAME=postgres
set ENV=local

docker container stop app-db

docker container run -d -p 5432:5432 --name app-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres postgres

timeout 1

go run adapters/ui/gofiber/*.go
