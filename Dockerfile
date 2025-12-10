FROM golang:1.25.5
WORKDIR /build
COPY *.go .
COPY go.mod .
RUN go build -o /app/myapp
EXPOSE 8081
CMD [ "/app/myapp" ] 