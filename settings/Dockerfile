FROM golang:1.20

WORKDIR /project

COPY cmd cmd/
COPY internal internal/
COPY go.mod go.sum ./

RUN go build -o ./bin/settings ./cmd/settings

ENTRYPOINT /project/bin/settings
