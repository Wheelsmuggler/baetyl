FROM --platform=$TARGETPLATFORM golang:1.13.5-stretch as devel
COPY / /go/src/
RUN cd /go/src/ && make build

FROM --platform=$TARGETPLATFORM busybox
COPY --from=devel /go/src/baetyl /bin/
ENTRYPOINT ["baetyl"]