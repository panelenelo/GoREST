FROM golang:1.26.2-alpine AS builder
WORKDIR /app
COPY . /app
RUN go build -o GoREST ./cmd/web
#CMD ["./GoREST"]

FROM scratch
WORKDIR /app
COPY --from=builder /app .
EXPOSE 8181
ENTRYPOINT []
CMD ["./GoREST"]
