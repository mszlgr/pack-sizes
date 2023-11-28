FROM golang:1.21 as builder
WORKDIR /build
COPY . .
RUN make static

FROM scratch
COPY --from=builder /build/pack-sizes /pack-sizes
CMD ["/pack-sizes"]
