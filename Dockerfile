FROM golang:1.19.3-alpine

WORKDIR /app
EXPOSE 9191
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/wishr-api
CMD ["wishr-api"]