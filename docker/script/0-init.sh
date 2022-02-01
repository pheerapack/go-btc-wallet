#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
  CREATE USER $APP_DB_USER WITH PASSWORD '$APP_DB_PASS';
  CREATE DATABASE $APP_DB_NAME;
  GRANT ALL PRIVILEGES ON DATABASE $APP_DB_NAME TO $APP_DB_USER;
  \connect $APP_DB_NAME $APP_DB_USER
  BEGIN;
    DROP TABLE IF EXISTS my_wallet;
    CREATE TABLE my_wallet
    (
      date_time TIMESTAMP (0) WITH TIME ZONE,
      amount numeric,
      PRIMARY KEY (date_time)
    );
    DROP TABLE IF EXISTS summary_by_hour;
    CREATE TABLE summary_by_hour
    (
      date_time TIMESTAMP (0) WITH TIME ZONE,
      amount numeric,
      PRIMARY KEY (date_time)
    );
  COMMIT;
EOSQL