FROM golang:1.25.5 AS builder

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -tags netgo,opusergo -a -trimpath -ldflags='-w -extldflags "-static" -buildid=' -o main .

FROM ghcr.io/greboid/dockerbase/nonroot:1.20251213.0

COPY --from=builder /app/main /httpredirect
CMD ["/httpredirect"]
