# Dockerfile for installing yakv

# Build the binary
FROM golang:1.14 as build
COPY . /src
WORKDIR /src
RUN go build -o vtec

# Add the Alpine Linux image
FROM alpine

# Copy binary
COPY --from=build /src/vtec .

# Copy certificate and key
COPY --from=build /src/*.pem .

EXPOSE 8081
