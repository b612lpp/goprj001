FROM golang:1.25.5 AS builder
WORKDIR /build
COPY . .
COPY go.mod .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/myapp.goprj001

FROM scratch
WORKDIR /
COPY --from=builder /app /app
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8081
CMD [ "/app/myapp.goprj001" ]