FROM golang:1.14-alpine AS build

WORKDIR /src/ms-product/

COPY ./ /src/ms-product/

EXPOSE 8080
 
RUN CGO_ENABLED=0 go build -o /bin/ms-product

FROM scratch
COPY --from=build /bin/ms-product /bin/ms-product
COPY --from=busybox:1.28 /bin/busybox /bin/busybox
ENTRYPOINT ["/bin/ms-product"]