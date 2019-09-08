FROM alpine:3.10.2

COPY modem-scraper /modem-scraper

VOLUME [ "/config" ]

ENTRYPOINT [ "./modem-scraper", "-config", "/config/config.yaml" ]
