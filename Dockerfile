FROM golang:1.17-buster as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . . 
RUN make build

FROM scratch 
COPY --from=builder /app/server /go/bin/hello
EXPOSE 8080
ENTRYPOINT ["go/bin/hello"]