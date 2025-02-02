FROM golang:alpine

WORKDIR /api

COPY . .

RUN go mod init github.com/ramisoul84/reset-server && go mod tidy

RUN go build -o main ./cmd/reset-server

EXPOSE 8080

CMD [ "./main" ]
