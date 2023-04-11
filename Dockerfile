FROM ghcr.io/greboid/dockerfiles/golang@sha256:b39e962ca9b7c2d31ba231c4912fc7831d59dfbb5dcd5e3fa9bba79bd51cc32c as builder

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -trimpath -ldflags=-buildid= -o main .

FROM ghcr.io/greboid/dockerfiles/base@sha256:ef8b65fda5d776551a87f970b02ebd1aa67f98c60b68637d910429a4e18be9b9

COPY --from=builder /app/main /httpredirect
EXPOSE 8080
CMD ["/httpredirect"]
