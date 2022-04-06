FROM golang:1.16-alpine as builder

ENV CGO_ENABLED=0

WORKDIR /build
COPY . .

# Generate code
RUN go get github.com/google/wire/cmd/wire && ./scripts/wire_generate.sh

# Build
RUN go build -o app cmd/todo-server/main.go

FROM alpine:3.11.6
COPY --from=builder /build/app .

CMD "./app"
