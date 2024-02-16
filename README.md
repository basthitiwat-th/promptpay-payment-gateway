# PROMPTPAY-PAYMENT_GATEWAY

1. start docker mysql
   `docker-compose -f compose-mysql-3306.yaml up -d`

2. use script `folder/script` for create table.

3. open new terminal use command `make start-promptpay`

4. open new terminal use command `make start-merchant`

5. open new terminal use command `make start-payment-gateway`

6. use curl.http in folder `test` for test.
