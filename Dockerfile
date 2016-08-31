FROM scratch
MAINTAINER https://github.com/damianodonati

COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ADD gobin /app/gobin
CMD ["/app/gobin"]
