# golang-study

## example booking app

    cd booking-app
    go run .

## the restapi

the go.mod, go.sum need to be in the root folder of vscode so that the lint server will know where the github packages are
go build restapi/main.go

    docker build -t yijunx/golang-restapi:0.0.1 .
    docker run -it --rm -p 8888:8001 yijunx/golang-restapi:0.0.1
