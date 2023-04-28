FROM golang:1.19.7-alpine3.17 AS srv-dependencies
ENV NAME "srv"
WORKDIR /opt/demomon/app
RUN ls -la
COPY go.mod .
COPY go.sum .


FROM srv-dependencies AS build
ENV NAME "srv"
WORKDIR /opt/demomon/app
COPY . .
RUN go build cmd/app/app.go

CMD ./app

