FROM golang:1.16 as builder

WORKDIR /build

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
# `skaffold debug` sets SKAFFOLD_GO_GCFLAGS to disable compiler optimizations
ARG SKAFFOLD_GO_GCFLAGS
# CGO has to be disabled for alpine
ENV CGO_ENABLED=0
RUN go build -gcflags="${SKAFFOLD_GO_GCFLAGS}" -o /app

FROM alpine:3.10
RUN apk --no-cache add tzdata
WORKDIR /
ENV GOTRACEBACK=single
COPY --from=builder /app /app
COPY --from=builder /build/config.json /config.json
COPY --from=builder /build/migrations /migrations
ENTRYPOINT ["/app"]
