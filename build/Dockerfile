FROM golang:1.22-alpine3.19 AS builder

COPY app /github.com/kms-qwe/microservices_course_auth/source/
WORKDIR /github.com/kms-qwe/microservices_course_auth/source/

RUN go mod download
RUN go build -o ./bin/crud_server ./cmd/grpc_server/main.go

FROM alpine:3.19

WORKDIR /root/
COPY --from=builder /github.com/kms-qwe/microservices_course_auth/source/bin/crud_server .
COPY config ./config/

CMD sh -c "./crud_server -config-path=${CONFIG_PATH}"




