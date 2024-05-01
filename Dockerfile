FROM golang:1.22-alpine

# Get git
RUN apk update && apk add --no-cache git

# Move to working directory
WORKDIR /app

# Copy App to workdir
COPY . .

# Download dependency
RUN go mod tidy

#  Build App
RUN go build -o binary

EXPOSE 8080

ENTRYPOINT [ "/app/binary" ]