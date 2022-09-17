# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY . .

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed.
RUN go mod download
# Ensures that the go.mod file matches the source code in the module.
RUN go mod tidy
# rerun a command when certain files change.
# https://github.com/cespare/reflex
RUN go install github.com/cespare/reflex@latest