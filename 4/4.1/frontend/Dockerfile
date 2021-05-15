FROM golang:alpine AS build
WORKDIR /build
COPY . .
RUN go mod download && \
    go build -o /out/main ./cmd

FROM alpine
COPY --from=build /out/main /build/templates ./
CMD ["./main"]
