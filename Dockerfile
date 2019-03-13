# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM gobuffalo/buffalo:v0.14.0 as builder

RUN mkdir -p $GOPATH/src/worldlocations
WORKDIR $GOPATH/src/worldlocations

RUN mkdir /app_src
ADD . .
RUN go get $(go list ./... | grep -v /vendor/)
RUN buffalo build --static -o /app_src/app

COPY docker-entrypoint.sh /app_src/docker-entrypoint.sh
RUN chmod +x /app_src/docker-entrypoint.sh

COPY secrets/wlio-sqlproxy-sa.json /app_src/wlio-sqlproxy-sa.json

#--------------------------
#    SQLPROXY
#--------------------------
RUN wget https://dl.google.com/cloudsql/cloud_sql_proxy.linux.amd64 -O /app_src/cloud_sql_proxy \
    && chmod +x /app_src/cloud_sql_proxy

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache wget

WORKDIR /app_src

COPY --from=builder /app_src .

# Uncomment to run the binary in "production" mode:
ARG env_arg=production
ENV GO_ENV=$env_arg

# Bind the app to 0.0.0.0 so it can be seen from outside the container
ENV ADDR=0.0.0.0

EXPOSE 8080

# Uncomment to run the migrations before running the binary:
# CMD /bin/app migrate; /bin/app

ENTRYPOINT ["/app_src/docker-entrypoint.sh"]