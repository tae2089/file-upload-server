FROM golang:1.19.3-alpine as build

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/main

FROM scratch
COPY --from=build /bin/main /bin/main
ENTRYPOINT ["/bin/main"]