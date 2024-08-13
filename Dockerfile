FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app/

# Add this line to inspect the directory structure
RUN ls -R /app

RUN go build -o vendor-service ./cmd/vendor-service

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/vendor-service .

CMD [ "./vendor-service" ]
