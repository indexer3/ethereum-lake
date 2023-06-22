FROM golang:1.20.4-alpine AS builder

WORKDIR /root/ethereum-lake

COPY . .

RUN make build 

FROM alpine:3.18.0 AS runner 

WORKDIR /root/ethereum-lake

RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /root/ethereum-lake/build/ethereum-lake .

ENTRYPOINT [ "ethereum-lake" ]

