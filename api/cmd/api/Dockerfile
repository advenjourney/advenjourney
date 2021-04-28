# build stage
FROM golang:1.15.7 AS builder
COPY . /src
WORKDIR /src/api
RUN make clean && make build

# final stage
FROM alpine:3.13
RUN apk add --no-cache ca-certificates
WORKDIR app
COPY --from=builder /src/api/db /app/db
COPY --from=builder /src/api/bin/api /app/api
ENTRYPOINT ["/app/api"]
