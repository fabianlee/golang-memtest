# multi-stage build so that:
#    golang builder is not needed on host
#    golang builder remnants not required in Docker image


#
# builder image
#
FROM golang:1.13-alpine3.11 as builder
RUN mkdir /build
ADD *.go /build/
WORKDIR /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o golang-memtest .


#
# generate clean, final image for end users
#
FROM alpine:3.11.3

# copy golang binary into container
COPY --from=builder /build/golang-memtest .

# executable
ENTRYPOINT [ "./golang-memtest" ]
# arguments that can be overridden
# 3Mb, 300 milliseconds between allocation
CMD [ "3", "300" ]

