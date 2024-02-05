# build stage
FROM golang:1.21-alpine AS build

ARG ENV_FILE
ARG FIREBASE_PATH

WORKDIR /app

COPY ./go.mod ./go.sum ./

COPY ./config/dev.env ./.env

COPY ./code-with-azri-firebase-adminsdk-vf6lc-bd1eeda135.json ./firebase-credentials.json

RUN go mod download

COPY ./ ./

RUN go build -o ./app ./cmd/main.go

# Run the tests in the container
FROM build AS run-test
RUN go test -v ./...

# deploy stage
FROM golang:1.21-alpine

WORKDIR /app

COPY --from=build ./app ./

EXPOSE 8080

ENTRYPOINT ["./app"]
