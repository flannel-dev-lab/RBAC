FROM mysql:5.6

ENV MYSQL_DATABASE=rbac \
    MYSQL_ROOT_PASSWORD=toor \
    MYSQL_USER=root

ADD database/mysql/schema/init.sql /docker-entrypoint-initdb.d
EXPOSE 3306