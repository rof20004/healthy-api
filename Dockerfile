FROM golang:latest as build-app

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o app-binary adapters/ui/gofiber/*.go

FROM gcr.io/distroless/static-debian11

COPY migrations ./migrations
COPY --from=build-app /app/app-binary ./app-binary

EXPOSE 8080

CMD ["./app-binary"]
