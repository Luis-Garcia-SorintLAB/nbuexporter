FROM alpine:3.17 AS builder
COPY /home/runner/work/nbuexporter/nbuexporter/nbuexporter .
RUN ./nbuexporter
CMD ["/bin/sh"]
