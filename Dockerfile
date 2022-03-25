# build layer
FROM golang:1.17-alpine AS build
RUN apk add build-base git
WORKDIR /app
COPY . .

RUN go mod download
RUN go build .



# app layer
FROM alpine:latest
WORKDIR /app
EXPOSE 8080

ENV APP=linkgo
ENV UID=1000
ENV GID=1000
RUN addgroup -g $GID $APP &&\
    adduser \
    --disabled-password \
    --gecos "" \
    --home "${pwd}" \
    --ingroup "$APP" \
    --no-create-home \
    --uid "$UID" \
    "$APP"
ENV TZ="Asia/Shanghai"
RUN apk add --no-cache tzdata

COPY --from=build --chown=$APP:$APP /app/${APP} /app/${APP}
# COPY --from=build --chown=$USER:$GROUP /app/templates /app/templates

# RUN mkdir /data && chown ${USER}:${GROUP} /data
CMD [ "/app/linkgo" ]