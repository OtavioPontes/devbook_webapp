FROM golang:alpine

WORKDIR /webapp

COPY . .  

RUN go build

EXPOSE 3000

CMD ["./devbook_webapp"]
