# First stage
FROM golang:alpine AS build
WORKDIR /app
COPY go.mod .
COPY main.go .
RUN go mod tidy
RUN go build -o main

# Second stage
FROM alpine:latest 
COPY --from=build /app/main .
CMD ["./main"]


