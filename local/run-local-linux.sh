#!/bin/bash

# Set environment variables
export DATABASE_HOST=127.0.0.1
export DATABASE_PORT=5432
export DATABASE_USER=postgres
export DATABASE_PASS=postgres
export DATABASE_NAME=postgres

# Remove old containers
docker container stop app-db

# Create new containers
docker container run -d -p 5432:5432 --rm --name app-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres postgres

# Waiting containers start
sleep 1

go run ui/presentation/gofiber/*.go
