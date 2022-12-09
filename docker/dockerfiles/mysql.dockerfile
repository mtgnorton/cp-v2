FROM --platform=linux/amd64 mysql:5.7

MAINTAINER mtg

COPY ../sql/mysqld.cnf /etc/mysql/mysql.conf.d/mysqld.cnf

COPY  ../sql/structure_and_data.sql   /docker-entrypoint-initdb.d/structure_and_data.sql
