FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /what-fridge

EXPOSE 8080

CMD [ "/what-fridge" ]
