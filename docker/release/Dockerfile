FROM golang:1.18-alpine as builder

ENV CGO_ENABLED=0

WORKDIR /build
COPY . .

# Generate code
RUN go install github.com/google/wire/cmd/wire@latest && ./scripts/wire_generate.sh
RUN go install github.com/swaggo/swag/cmd/swag@latest && ./scripts/generate_swagger.sh

# Build
RUN go build -o app cmd/todo-server/main.go

FROM alpine:latest
COPY --from=builder /build/app .

CMD "./app"
