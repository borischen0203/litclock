# build stage
FROM golang:1.16-alpine AS build
ADD . /src
RUN cd /src && go build -o litclock

# final stage
FROM alpine as runtime

# add TimeZone tzdata
RUN apk add --no-cache tzdata
ENV TZ=Europe/London

WORKDIR /app
COPY --from=build /src/litclock /app/litclock
COPY run.sh /app/
RUN cd /app && chmod +x run.sh

ENTRYPOINT ["/app/run.sh"]

