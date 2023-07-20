FROM golang:1.20 as builder

ARG ATLAS_URI

ENV APP_HOME /go/src/exam
ENV ATLAS_URI ${ATLAS_URI}
ENV PORT 8080

RUN echo ${ATLAS_URI}

WORKDIR "$APP_HOME"

COPY . .

RUN go mod download
RUN go build -o exam

# copy build to a clean image
FROM golang:1.20

ARG ATLAS_URI

ENV APP_HOME /go/src/exam
ENV PORT 8080
ENV ATLAS_URI ${ATLAS_URI}

RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/exam $APP_HOME

EXPOSE $PORT

CMD ["./exam"]
