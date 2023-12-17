FROM golang:1.20.6

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /application

EXPOSE 8000

CMD [ "/application" ]