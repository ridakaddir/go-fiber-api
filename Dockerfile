FROM golang:1.18 as builder

WORKDIR /app

COPY . .

RUN echo $(ls -1 /app)

RUN CGO_ENABLED=0 go build -o main main.go

RUN echo $(ls -1 /app)

FROM alpine:3.15

WORKDIR /app

COPY --from=builder /app/main .

RUN echo $(ls -1 /app)

EXPOSE 3000

CMD ["./main"]
