FROM golang:alpine AS build
WORKDIR /src
COPY . .
RUN go get -d -v && \
    go build -o /out/main

FROM scratch
COPY --from=build /out/main .
ENTRYPOINT ["./main"]
