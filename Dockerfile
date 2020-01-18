FROM alpine:latest

COPY golang-memtest .

ARG nmb=1
ENV env_nmb=$nmb

CMD [ "./golang-memtest" ]
