FROM golang:alpine

WORKDIR /pass-manager

COPY go.mod go.sum ./
RUN go mod tidy

COPY . ./


RUN CGO_ENABLED=0 GOOS=linux go install .

WORKDIR /

CMD [ "bin/sh" ]
