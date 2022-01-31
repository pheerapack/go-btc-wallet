\c aggregation

BEGIN;
DROP TABLE IF EXISTS summary_by_hour;
CREATE TABLE summary_by_hour
(
   date_time TIMESTAMP (0) WITH TIME ZONE,
   amount numeric,
   PRIMARY KEY (date_time)
);
END;