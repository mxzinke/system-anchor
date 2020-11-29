FROM golang:1.15.2
WORKDIR /app

COPY . .

RUN go get
RUN go build -o system-anchor .

EXPOSE 80
ENTRYPOINT ["./system-anchor"]