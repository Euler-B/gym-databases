FROM mysql:latest

#Mantainer info
LABEL mantainer="Eduardo Bravo <eduardojosebb.matescience@gmail.com>"

RUN chown -R mysql:root /var/lib/mysql    

ARG MYSQL_DATABASE
ARG MYSQL_USER
ARG MYSQL_PASSWORD
ARG MYSQL_PASSWORD_ROOT

ENV MYSQL_DATABASE=${MYSQL_DATABASE}
ENV MYSQL_USER=${MYSQL_USER}
ENV MYSQL_PASSWORD=${MYSQL_PASSWORD}
ENV MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD}

ADD create-tables.sql /etc/mysql/create-tables.sql

RUN sed -i 's/MYSQL_DATABASE/'$MYSQL_DATABASE'/g' /etc/mysql/create-tables.sql
RUN cp /etc/mysql/create-tables.sql /docker-entrypoint-initdb.d

EXPOSE 3306



