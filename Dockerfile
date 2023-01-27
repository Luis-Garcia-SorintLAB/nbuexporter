FROM alpine:3.17 AS builder
COPY nbuexporter .
RUN ./nbuexporter
CMD ["/bin/sh"]
