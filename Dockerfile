FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o birthday-bot ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/birthday-bot .

# 타임존 설정 (한국 시간)
RUN apk add --no-cache tzdata
ENV TZ=Asia/Seoul

CMD ["./birthday-bot"]