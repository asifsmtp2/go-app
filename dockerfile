FROM golang:1.22.1

WORKDIR /app

# Copying content from source to destination
COPY go.mod .
COPY main.go .

# To import third party dependencies
RUN go get

EXPOSE 8080

# To prepare build
RUN go build -o bin .

ENTRYPOINT [ "/app/bin" ]