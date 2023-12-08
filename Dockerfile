FROM node:15.14.0-alpine3.10 as node-builder
COPY ./web /src/web
WORKDIR /src/web
RUN yarn && yarn exec vuepress build docs

FROM golang:1.16.4 AS go-builder
COPY ./api /src/api
WORKDIR /src/api
RUN make clean && make sync
COPY --from=node-builder /src/web/docs/.vuepress/dist /src/api/assets
RUN make generate && make build

# final stage
FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR app
COPY --from=go-builder /src/api/db /app/db
COPY --from=go-builder /src/api/bin/api /app/api
ENTRYPOINT ["/app/api"]
