FROM alpine:3.17 AS builder
RUN mkdir output
RUN mkdir config
COPY nbuexporter .
COPY servers/* ./config
CMD ["sh nbuexporter"]
