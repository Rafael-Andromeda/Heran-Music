FROM golang:1.23-alpine AS build
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /heran-music ./cmd/api

FROM alpine:3.21
RUN apk --no-cache add ca-certificates
COPY --from=build /heran-music /heran-music

EXPOSE 8080
USER 65532:65532
ENTRYPOINT ["/heran-music"]
