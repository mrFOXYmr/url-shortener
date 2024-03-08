FROM postgres:12-bullseye
COPY init.sql /docker-entrypoint-initdb.d/

env POSTGRES_USER=psql
env POSTGRES_PASSWORD=psql
env POSTGRES_DATABASE=db
