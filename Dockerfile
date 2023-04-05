FROM postgres:10.17

ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=postgres
ENV POSTGRES_DB=jwt

COPY ./*.sql /docker-entrypoint-initdb.d/

EXPOSE 5432