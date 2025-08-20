# Etapa 1 - Build
FROM golang:1.24 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./main.go

# Etapa 2 - Runtime
FROM gcr.io/distroless/base-debian12
WORKDIR /app

# Copia o binário
COPY --from=builder /app/server .

# Copia o arquivo .env
COPY .env ./

EXPOSE 8080
CMD ["./server"]
