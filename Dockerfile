# multi-stage build so that:
#    golang builder is not needed on host
#    golang builder remnants not required in Docker image


# builder image
FROM golang:1.13-alpine3.11 as builder
RUN mkdir /build
ADD *.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o golang-memtest .
RUN find /build



# generate clean, final image for end users
FROM alpine:3.11.3

# copy golang binary into container
COPY --from=builder /build/golang-memtest .

# number of Mb to allocate, default 1Mb
ARG nmb=1
ENV env_nmb=$nmb

# number of Millseconds to wait between mem allocations, default 100ms
ARG nms=100
ENV env_nms=$nms

CMD [ "./golang-memtest" ]
