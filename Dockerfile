FROM golang:1.22

ENV GO111MODULE=on

RUN mkdir -p /microservice
ADD . /microservice
WORKDIR /microservice

RUN go mod download
COPY . .
RUN go build ./cmd/app/main.go

RUN chmod a+x ./scripts/start_service.sh
ENTRYPOINT ["./scripts/start_service.sh"]