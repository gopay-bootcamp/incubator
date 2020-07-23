# build generator binary
FROM golang:alpine AS builder

WORKDIR /app
ADD . .
RUN go build -o hello

FROM alpine

COPY --from=builder /app/hello hello

ENTRYPOINT [ "./hello" ]
