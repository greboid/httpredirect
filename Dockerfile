FROM golang:latest as builder
WORKDIR /app 
COPY main.go . 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o main .

FROM scratch
COPY --from=builder /app/main /main
EXPOSE 80
CMD ["/main"]