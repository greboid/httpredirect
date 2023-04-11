FROM ghcr.io/greboid/dockerfiles/golang@sha256:25cefe2da86b16981762f77a2d8b1ed2b611e62d92e379dd78d05c3211ca4a21 as builder

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .

FROM ghcr.io/greboid/dockerfiles/base@sha256:82873fbcddc94e3cf77fdfe36765391b6e6049701623a62c2a23248d2a42b1cf

COPY --from=builder /app/main /httpredirect
EXPOSE 8080
CMD ["/httpredirect"]
