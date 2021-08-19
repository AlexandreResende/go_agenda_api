FROM golang:1.16

WORKDIR /app

# the . represents the workdir - in this case /app
COPY go.mod .
COPY go.sum .

# this will download all dependencies
RUN go mod download

COPY . .

#ENV port 8000

RUN go build

CMD ["./go_agenda_api"]

# docker build -t go_agenda_api
# docker run -p 8000:8000 go_agenda_api