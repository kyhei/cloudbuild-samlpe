FROM golang:1.15-alpine as builder

WORKDIR /opt/build

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o app


FROM alpine:3.12
ENV PORT 8080

COPY --from=builder /opt/build/app /opt/app

WORKDIR /opt
CMD ["./app"]