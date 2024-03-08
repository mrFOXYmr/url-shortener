from golang:alpine as builder
workdir /app
copy ./ ./
RUN go mod download
RUN go build main.go

from alpine
workdir /app
copy --from=builder /app/main /app/main
ENTRYPOINT ["./main"]
