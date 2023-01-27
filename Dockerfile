FROM alpine:3.17 AS builder
RUN mkdir ./output
COPY nbuexporter .
COPY servers/* ./output
CMD ["./nbuexporter"]
