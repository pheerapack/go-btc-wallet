\c aggregation

BEGIN;
DROP TABLE IF EXISTS my_wallet;
CREATE TABLE my_wallet
(
   date_time TIMESTAMP (0) WITH TIME ZONE,
   amount numeric,
   PRIMARY KEY (date_time)
);
END;