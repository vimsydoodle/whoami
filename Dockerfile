FROM golang:alpine3.15 AS builder
ADD . /app
WORKDIR /app
RUN go build -o http

FROM alpine:3.15
WORKDIR /app
ENV PORT 8000
EXPOSE 8000
COPY --from=builder /app/http /app
CMD ["/app/http"]
