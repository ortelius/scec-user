FROM cgr.dev/chainguard/go@sha256:74bc9af1d45fd1c8d432a89148c5e413711204636b54ca05197b511bea7a18fb AS builder

WORKDIR /app
COPY . /app

RUN go mod tidy; \
    go build -o main .

FROM cgr.dev/chainguard/glibc-dynamic@sha256:6462e815d213b0f339e40f84cbd696a9e4141ef4d62245a4ba5378cf09801527

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs docs

ENV ARANGO_HOST localhost
ENV ARANGO_USER root
ENV ARANGO_PASS rootpassword
ENV ARANGO_PORT 8529
ENV MS_PORT 8080

EXPOSE 8080

ENTRYPOINT [ "/app/main" ]
