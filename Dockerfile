#
#FROM golang:1.24-alpine
#WORKDIR /app
#
#
#RUN apk add --no-cache git ca-certificates && update-ca-certificates
#
#ENV GOPROXY=https://goproxy.io,direct
#
#
#COPY go.mod go.sum ./
#RUN go mod download
#
#
#COPY . .
#RUN go build -o main .
#
#EXPOSE 8080
#
#CMD ["./main"]
#
