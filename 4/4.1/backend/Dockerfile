FROM golang:alpine AS build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /out/main ./cmd

FROM alpine
COPY --from=build /out/main .
CMD ["./main"]
