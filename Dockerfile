FROM golang:1.20 AS build
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest as prodection
COPY --from=build /app .
CMD [ "./app" ]