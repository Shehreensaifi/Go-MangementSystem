#build stage
FROM --platform=linux/amd64 public.ecr.aws/i1l4e9j0/golang:1.19-alpine3.17
RUN apk add git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o main main.go
CMD ["./main"]