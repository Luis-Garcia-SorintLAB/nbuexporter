FROM gcc:11.3.0
RUN mkdir output
RUN mkdir config
COPY nbuexporter .
CMD ["./nbuexporter"]