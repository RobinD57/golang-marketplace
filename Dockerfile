FROM golang as build-stage
WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download

ADD pkg/chat /app/chat
COPY /cmd/web/main.go /app

RUN cd /app && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app
RUN ENV=${ENV}

FROM alpine
COPY --from=build-stage /app/app /
CMD ["/app"]
