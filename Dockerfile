FROM golang:1.25.3-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./
COPY pkg ./pkg

RUN CGO_ENABLED=0 GOOS=linux go build -o /dock

FROM alpine:3.19

COPY --from=build /dock /dock

CMD ["/dock"]

LABEL org.opencontainers.image.source="https://github.com/datasektionen/dock" \
      org.opencontainers.image.licenses="MIT"
