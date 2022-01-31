FROM postgres:13.1-alpine

COPY /docker/script/00_createDB.sh /docker-entrypoint-initdb.d/00_init.sh
COPY /docker/script/*.sql /docker-entrypoint-initdb.d/

RUN chmod +x /docker-entrypoint-initdb.d/00_init.sh

# Set host time zone 
ENV TZ=Asia/Bangkok
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone