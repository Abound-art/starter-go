FROM golang:1.18-alpine AS build
WORKDIR /build
RUN apk add upx git
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags "-s -w -extldflags '-static'" -o ./app ./lorenz/main
RUN upx ./app

FROM scratch
COPY --from=build /build/app /app

ENTRYPOINT ["/app"]