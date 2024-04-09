FROM golang:1.21.5-alpine3.19 AS builder
WORKDIR /go/src/shipping
COPY . . 
RUN go build -o shipping cmd/main.go

FROM scratch
COPY --from=builder /go/src/shipping/shipping .
COPY --from=builder /go/src/shipping/.env .
ENTRYPOINT [ "./shipping" ]