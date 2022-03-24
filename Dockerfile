FROM golang:1.18

WORKDIR /code

COPY . .

RUN go build restapi/main.go

# no need to EXPOSE 8000

ENTRYPOINT [ "./main" ]