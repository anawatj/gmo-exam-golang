FROM golang:1.18.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./app

EXPOSE 5000


CMD ["go","run","."]