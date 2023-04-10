FROM golang:alpine as base

ARG ARG_DB_HOST
ARG ARG_DB_PORT
ARG ARG_DB_USER
ARG ARG_DB_PASS

ENV DB_HOST=ARG_DB_HOST
ENV DB_PORT=ARG_DB_PORT
ENV DB_USER=ARG_DB_USER
ENV DB_PASS=ARG_DB_PASS

FROM base as builder
WORKDIR /go/src/api
COPY . .
RUN go build main.go

FROM base as runner
WORKDIR /app
COPY --from=builder /go/src/api /app
CMD go run main.go

EXPOSE 8080