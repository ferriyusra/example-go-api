#syntax=docker/dockerfile:1
FROM golang:1.22-alpine 

RUN apk update && apk upgrade && apk --update add git make

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0

COPY . .

RUN go build -o /academy

EXPOSE 8080

CMD [ "/academy" ]
