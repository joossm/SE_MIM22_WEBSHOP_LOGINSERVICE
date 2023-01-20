FROM golang:1.19

WORKDIR /SE_MIM22_WEBSHOP_LOGINSERVICE

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8441

RUN go mod download

ENTRYPOINT go build  && ./SE_MIM22_WEBSHOP_LOGINSERVICE