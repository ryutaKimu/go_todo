FROM golang:1.25.0-bookworm

WORKDIR /app

RUN go install github.com/air-verse/air@latest \
  && go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum  ./
RUN go mod tidy

COPY . .

RUN mkdir -p tmp

CMD ["air"]