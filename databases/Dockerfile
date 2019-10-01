# this image will serve as a base for the worldlocations.io
# it just includes the database
#
# tag gcr.io/worldlocation-io/worldlocations-db:[VERSION]
FROM gobuffalo/buffalo:v0.14.0
RUN mkdir -p /var/databases
COPY worldlocations_production.sqlite /var/databases/worldlocations_production.sqlite