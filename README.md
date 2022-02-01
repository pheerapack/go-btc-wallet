# go-btc-wallet

BTC-Wallet app
This application contail 2 type of service
1. Rest api for insert btc to wallet , inquiry summary btc eatch hour and update btc in wallet

    /store >> save btc history to wallet database (my_wallet table)and summary to each hour (summay_in_hour talbe)

    /inquiry >> inquiry between time (summay_in_hour table)

    /update  >> summary history to each hour (summay_in_hour table)

This is concept abount api will be use only when btc in/out but will not use so ofthen than /inquiry api , so all logic that calculate will be store in /store api.

2. Back ground process to update btc in wallet

This background process can start to update table summay_in_hour .
It is like monitor , sleep and wake up to summary depend on env.MONITORSLEEPTIME = number * time.Second

Require Tool before run :

- docker
- go version 1.17
- postman

How to run it :

Set up env in docker-compose.yml first

```
      - PG_HOSTNAME=yourlocalip
      - PG_HOSTPORT=5003
      - PG_USERNAME=postgres
      - PG_PASSWORD=postgres
      - PG_DBNAME=my_wallet
```

this is command that can help run this app , run it in order

`docker-compose build`

`docker-compose up db`

`docker-compose up btc_wallet_monitor`

`docker-compose up btc_wallet`

and this one for clean db

`docker-compose down`

Test case
Prerequisite data , insert into table

```
    INSERT INTO my_wallet(date_time, amount) VALUES ('2019-10-07T15:01:07+07:00', 10);
    INSERT INTO my_wallet(date_time, amount) VALUES ('2019-10-07T15:30:07+07:00', 1.1);
    INSERT INTO my_wallet(date_time, amount) VALUES ('2019-10-07T15:45:07+07:00', 7);
    INSERT INTO my_wallet(date_time, amount) VALUES ('2019-10-07T17:11:07+07:00', 100);
    INSERT INTO my_wallet(date_time, amount) VALUES ('2019-10-07T22:05:07+07:00', 10);
    INSERT INTO my_wallet(date_time, amount) VALUES ('2019-10-07T22:25:07+07:00', 10);
    INSERT INTO my_wallet(date_time, amount) VALUES ('2019-10-08T08:45:07+07:00', 10);
    INSERT INTO my_wallet(date_time, amount) VALUES ('2019-10-20T17:45:07+07:00', 10);
```

Add new btc to wallet by call api using postman and import file My-BTC-Wallet Test case.postman_collection.json

Run Folllowing this test case to test it

1.query all btc in wallet and summary by hour before insert

2.save btc to wallet (wait 10 sec , to monitor update)

3.query all btc in wallet and summary by hour after insert

4.query all btc in wallet and summary by invalid input

5.query all btc in wallet and summary by same hour

all test should be pass