FROM golang:1.21.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /reddit-server

EXPOSE 8080

# Run
CMD ["/reddit-server"]