FROM golang:alpine3.16 as builder

WORKDIR /app
RUN apk add --no-cache gcc musl-dev linux-headers git
COPY . /app/
RUN cd /app && go mod tidy
RUN cd /app && go build -o ./build/bin/main -gcflags="all=-N -l" ./cmd/main.go


FROM alpine:latest

WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/build/bin /app

ENTRYPOINT [ "/app/main" ]
