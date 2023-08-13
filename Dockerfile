FROM golang:1.19-alpine AS builder

RUN go version

ARG PROJECT_VERSION


WORKDIR /go/src/app
COPY . .
RUN set -Eeux && \
    go mod download && \
    go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -trimpath \
    -ldflags="-w -s -X 'main.Version=${PROJECT_VERSION}'" \
    -o app cmd/miniwiki/miniwiki.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /go/src/app/app .
COPY --from=builder /go/src/app/api/ ./api/

EXPOSE 8000

CMD [ "./app"]
