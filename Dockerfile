# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/

##
# NOTE
# this repo can be easily built with the
# FROM gobuffalo/buffalo:v0.14.0 as builder
# but the database will be missing
# I cannot share the database because I do not have the right to distribute
FROM gcr.io/worldlocation-io/worldlocations-db:v0.0.1 as builder

RUN mkdir -p $GOPATH/src/worldlocations
WORKDIR $GOPATH/src/worldlocations

RUN mkdir /app_src
ADD . .
RUN go get $(go list ./... | grep -v /vendor/)
RUN buffalo build --static -o /app_src/app

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache wget

WORKDIR /app_src

COPY --from=builder /app_src .

RUN mkdir -p /var/databases
COPY databases/worldlocations_production.sqlite /var/databases/worldlocations_production.sqlite

# Uncomment to run the binary in "production" mode:

# Bind the app to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0

EXPOSE 8080

# Uncomment to run the migrations before running the binary:
# CMD /bin/app migrate; /bin/app
CMD exec /app_src/app;