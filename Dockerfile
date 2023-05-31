#FROM golang:alpine AS builder
#WORKDIR /app
#COPY . .
#RUN go mod download
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o crud ./cmd/main.go
#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
#WORKDIR /root/
#COPY --from=builder /app/crud .
#CMD ["./crud"]

#FROM golang:alpine AS builder
#WORKDIR /app
#COPY . .
#RUN apk add build-base && go build -o forum cmd/main.go
#FROM alpine:latest
#WORKDIR /app
#COPY --from=builder /app .
#CMD ["./forum"]
