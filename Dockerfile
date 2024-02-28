FROM golang:1.22.0 as builder

WORKDIR /app

COPY . .

# build the application statically for linux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

# expects either http, grpc, or graphql - will default to http
ARG API=http
ENV API=${API}

WORKDIR /root/

COPY --from=builder /app/main .

# run the application with the given server
CMD ./main -api $API
