FROM golang:alpine as base

FROM base as builder
WORKDIR /go/src/api
COPY . .
RUN go build main.go

FROM base as runner
WORKDIR /app
COPY --from=builder /go/src/api/main .
CMD ./main

EXPOSE 8080