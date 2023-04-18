FROM golang

WORKDIR /go/src/app
COPY . .
# RUN go build -o main .
RUN go build
EXPOSE 8080
CMD ["./main"]