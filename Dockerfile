FROM golang:1.27.2 AS builder

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -tags netgo,opusergo -a -trimpath -ldflags='-s -w -extldflags "-static" -buildid=' -o main .

FROM ghcr.io/greboid/dockerbase/nonroot:1.20261020.1

COPY --from=builder /app/main /httpredirect
CMD ["/httpredirect"]
