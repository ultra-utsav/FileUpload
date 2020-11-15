FROM golang:latest

LABEL maintainer="Utsav <utsav.r.vaghani@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build

CMD ["./FileUpload.git.exe"]

CMD tail -f /dev/null