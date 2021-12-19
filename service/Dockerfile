##
## Build
##
FROM golang:1.17-bullseye AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /service

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /service /service

EXPOSE 8090

USER nonroot:nonroot

ENTRYPOINT ["/service"]
