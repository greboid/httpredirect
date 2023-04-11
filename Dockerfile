FROM ghcr.io/greboid/dockerfiles/golang:latest as builder

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .

FROM ghcr.io/greboid/dockerfiles/base:latest

COPY --from=builder /app/main /httpredirect
EXPOSE 8080
CMD ["/httpredirect"]
