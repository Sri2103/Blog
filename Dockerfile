FROM  golang:1.21.6 as builder
WORKDIR  /app
COPY go.mod   .
COPY go.sum   .
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.15 AS final
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 4500
CMD ["./main"]
