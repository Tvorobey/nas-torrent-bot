from alpine

COPY torrent_bot .
RUN apk add --no-cache ca-certificates

EXPOSE 8080

CMD ["/notifier"]