FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

EXPOSE 8191

# Usar go run directamente
CMD ["go", "run", "main.go"]