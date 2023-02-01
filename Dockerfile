FROM golang:1.19-alpine AS build

WORKDIR /build

RUN apk add upx
COPY go.mod ./
COPY go.sum ./
COPY abound/ ./abound/
COPY cmd/ ./cmd/
COPY algo/ ./algo/

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w -extldflags '-static'" -o app ./cmd/algo
RUN upx app

FROM scratch
COPY --from=build /build/app /app

ENTRYPOINT ["/app"]