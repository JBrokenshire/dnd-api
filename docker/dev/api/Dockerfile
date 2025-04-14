# Start from golang base image
FROM golang:1.23-alpine as builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

RUN cd /app && go install github.com/githubnemo/CompileDaemon@latest
RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3 && go install github.com/go-swagger/go-swagger/cmd/swagger@v0.31.0

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

# Command to run the executable
CMD  swag init -g cmd/api/api.go \
  && swagger validate docs/swagger.yaml \
  && /wait \
  && go run db/migrations/entry.go --verbose \
  && CompileDaemon --build="go build -o /app/cmd/api/api /app/cmd/api/api.go"  --command="/app/cmd/api/api" --color
