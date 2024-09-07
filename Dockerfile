FROM docker.io/library/golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build github.com/basado1991/jwt_auth_service/cmd/auth_service

CMD ["./auth_service"]
