# #  syntax=docker/dockerfile:1

# # ##
# # ## Build
# # ##
# # FROM golang:1.17-alpine

# # WORKDIR /app

# # COPY go.mod ./
# # COPY go.sum ./
# # RUN go mod download

# # COPY gochan.go ./

# # RUN go build -o /docker-chan

# # ##
# # ## Deploy
# # ##  multistage
# # # FROM gcr.io/distroless/base-debian10

# # # WORKDIR /

# # # COPY --from=build /docker-gs-ping /docker-gs-ping

# # # EXPOSE 8080

# # # USER nonroot:nonroot

# # ENTRYPOINT ["/docker-gs-ping"]
# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY gochan.go ./
# COPY index.html ./
RUN go build -o /docker-socket

EXPOSE 8080

CMD [ "/docker-socket" ]
# # syntax=docker/dockerfile:1

# FROM golang:1.9-alpine3.8

# WORKDIR /app

# # COPY go.mod ./
# # COPY go.sum ./
# RUN git clone github.com/googollee/go-socket.io

# COPY socket.go ./

# RUN go run socket.go

# # EXPOSE 8080

# # CMD [ "/docker-gs-ping" ]