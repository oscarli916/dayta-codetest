# Build first image & create executable

FROM golang:alpine AS build

WORKDIR /app

ADD . .

RUN cd client && go build .

# Prod build

FROM alpine

ENV SECRET_KEY = "SECRET_KEY"

WORKDIR /app

# TODO:
#   Finish deployment docker build & serve with an entrypoint
COPY --from=build /app/client /app/
EXPOSE 8080
ENTRYPOINT [ "./client" ]