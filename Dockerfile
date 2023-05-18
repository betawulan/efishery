FROM golang:1.17-alpine as build

WORKDIR /app

COPY . .

# download depedencies
RUN go mod vendor

# build binary
RUN go build -o app main.go

EXPOSE 5050

CMD ["./app"]