FROM alpine
COPY service .
ENTRYPOINT ["./service"]