FROM golang:1.23.4 AS builder
COPY *.go ./
# install dependencies
ENV GOPROXY=direct
COPY go.mod go.sum ./
RUN go mod download
# copy source code
COPY . .
# compile application
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

EXPOSE 8080

CMD ["./main"]