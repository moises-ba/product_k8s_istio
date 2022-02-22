FROM golang:1.14-alpine AS build

WORKDIR /src/ms-product-price/

COPY ./ /src/ms-product-price/

EXPOSE 8081
 
RUN CGO_ENABLED=0 go build -o /bin/ms-product-price

FROM scratch
COPY --from=build /bin/ms-product-price /bin/ms-product-price
COPY --from=busybox:1.28 /bin/busybox /bin/busybox
ENTRYPOINT ["/bin/ms-product-price"]