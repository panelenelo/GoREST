FROM golang:1.26.2-alpine AS builder
WORKDIR /app
RUN apk add --no-cache make gcc musl-dev linux-headers 
COPY . /app
RUN make
RUN apk del make gcc musl-dev linux-headers

FROM scratch
WORKDIR /app
COPY --from=builder /app /app
EXPOSE 8181
ENTRYPOINT []
CMD ["./GoRest"]
