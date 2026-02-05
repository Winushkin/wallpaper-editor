FROM golang:1.25-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o dog-server

FROM alpine
WORKDIR /app
COPY --from=build /app/dog-server .
EXPOSE 8080
CMD ["./dog-server"]
