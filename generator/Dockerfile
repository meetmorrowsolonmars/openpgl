FROM golang:1.20

WORKDIR /project

COPY cmd cmd/
COPY internal internal/
COPY go.mod go.sum ./

RUN go build -o ./bin/diamond_square ./cmd/diamond_square

ENTRYPOINT /project/bin/diamond_square
