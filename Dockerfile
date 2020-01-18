FROM alpine:latest

# copy golang binary into container
COPY golang-memtest .

# number of Mb to allocate, default 1Mb
ARG nmb=1
ENV env_nmb=$nmb

# number of Millseconds to wait between mem allocations, default 100ms
ARG nms=100
ENV env_nms=$nms

CMD [ "./golang-memtest" ]
