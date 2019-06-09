FROM golang:1.12 as build
WORKDIR /go/src/app
COPY . .
ENV GO111MODULE on
RUN go build -v -o /app .
# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /app /app
CMD ["/app"]