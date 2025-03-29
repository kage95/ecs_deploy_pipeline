FROM golang:1.23.4-alpine AS build
WORKDIR /app
COPY ./api /app
RUN GOOS=linux GOARCH=amd64 go build -o main main.go

FROM ubuntu:latest
COPY --from=build /app/main .
CMD [ "./main" ]
