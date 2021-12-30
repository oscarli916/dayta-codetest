# Build first image & create executable

FROM golang:alpine AS build

WORKDIR /app

ADD . .

RUN cd server && go build .

# Prod build

FROM alpine

ENV SECRET_KEY = "SECRET_KEY"

WORKDIR /app

# TODO:
#   Finish deployment docker build & serve with an entrypoint
EXPOSE 3030
WORKDIR [./server/server]