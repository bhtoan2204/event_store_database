FROM golang:1.24-alpine AS build

WORKDIR /usr/app

COPY app/event_sourcing_gateway/go.mod .
COPY app/event_sourcing_gateway/go.sum .
RUN go mod download

COPY app/event_sourcing_gateway/ .

RUN go build -o /usr/app/event_sourcing_gateway /usr/app/cmd/main.go

FROM alpine AS production

WORKDIR /usr/app

COPY --from=build /usr/app ./

CMD ["./event_sourcing_gateway"]
