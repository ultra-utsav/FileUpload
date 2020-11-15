FROM golang:latest

LABEL maintainer="Utsav <utsav.r.vaghani@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]

CMD tail -f /dev/null

