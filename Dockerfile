FROM alpine:3.17 AS builder
COPY nbuexporter .
CMD ["./nbuexporter"]
