FROM golang:1.17-buster as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . . 
RUN make build

FROM scratch AS bin
COPY --from=builder /app/lana-sre-challenge-carlos lana-sre-challenge-carlos
EXPOSE 8080
ENTRYPOINT [ "/lana-sre-challenge-carlos" ]