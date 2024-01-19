FROM golang:1.21.4

ARG ATLAS_URI
ARG FIREBASE_PROJECT_ID

ENV APP_HOME /go/src/authv2
ENV ATLAS_URI ${ATLAS_URI}
ENV FIREBASE_PROJECT_ID ${FIREBASE_PROJECT_ID}
ENV GIN_MODE=release

WORKDIR "$APP_HOME"

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -tags=jsoniter -o authv2 

EXPOSE 8080

CMD ["./authv2"]
