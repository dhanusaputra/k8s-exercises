FROM golang:alpine AS build
WORKDIR /src
COPY . .
RUN go get -d -v && \
    go build -o /out/main

FROM alpine
COPY --from=build /out/main .
CMD ["./main"]
