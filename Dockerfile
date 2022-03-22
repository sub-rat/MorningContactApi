FROM golang:1.18.0
LABEL author subrat
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./bin/contact cmd/main.go
CMD ["./bin/contact"]