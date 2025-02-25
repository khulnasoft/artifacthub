# Build backend
FROM golang:1.23.3-alpine3.20 AS backend-builder
WORKDIR /go/src/github.com/khulnasoft/artifacthub
COPY go.* ./
COPY cmd/artifacthub cmd/artifacthub
COPY internal internal
WORKDIR /go/src/github.com/khulnasoft/artifacthub/cmd/artifacthub
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /artifacthub .

# Build frontend
FROM node:23-alpine3.20 AS frontend-builder
RUN apk --no-cache add jq
WORKDIR /web
COPY web .
ENV NODE_OPTIONS=--max_old_space_size=4096
RUN yarn install
RUN yarn build

# Build widget
FROM node:23-alpine3.20 AS widget-builder
WORKDIR /widget
COPY widget .
ENV NODE_OPTIONS=--max_old_space_size=4096
RUN yarn install
RUN yarn build

# Build docs
FROM klakegg/hugo:0.111.3 AS docs-builder
WORKDIR /
COPY scripts scripts
COPY docs docs
RUN scripts/prepare-docs.sh
WORKDIR /docs/www
RUN hugo

# Final stage
FROM alpine:3.20.3
RUN apk --no-cache add ca-certificates && addgroup -S hub -g 1000 && adduser -S hub -u 1000 -G hub
USER 1000
WORKDIR /home/artifacthub
COPY --from=backend-builder /artifacthub ./
COPY --from=frontend-builder /web/build ./web
COPY --from=frontend-builder /web/package.json ./web
COPY --from=widget-builder /widget/build ./widget
COPY --from=docs-builder /web/build/docs ./web/docs
COPY --from=docs-builder /docs/api ./web/docs/api
CMD ["./artifacthub"]
EXPOSE 8000
