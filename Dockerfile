FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.* ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o /server .

FROM alpine:3.22

RUN apk add --no-cache ca-certificates tzdata wget

WORKDIR /app

COPY --from=builder /server /app/server

EXPOSE 8080

ENTRYPOINT ["/app/server"]
