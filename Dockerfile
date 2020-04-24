FROM golang:latest as builder

MAINTAINER huiming.zhang@cerence.com

ARG git_url=github.com/yxzhm/
ARG componet_name=hydra-login-consent-poc


COPY ./ /go/src/${git_url}${componet_name}
WORKDIR /go/src/${git_url}${componet_name}

ENV GO111MODULE=on
ENV CGO_ENABLED=0


RUN go build -i -o /go/bin/app ${git_url}${componet_name}

FROM alpine:3.10
COPY --from=builder /go/bin/app /
RUN chmod 777 /app
CMD   ["/app"]