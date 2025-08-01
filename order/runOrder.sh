#!/bin/bash
docker run -p 3306:3306 \
 -e MYSQL_ROOT_PASSWORD=verysecretpass \
 -e MYSQL_DATABASE=order mysql

export DATA_SOURCE_URL="root:verysecretpass@tcp(127.0.0.1:3306)/order"
export APPLICATION_PORT=3000 
export ENV=development
export PAYMENT_SERVICE_URL=localhost:3001
go run cmd/main.go

grpcurl -d '{"user_id": 123, "order_items": [{"product_code": "prod", "quantity": 4, "unit_price": 12}]}' -plaintext localhost:3000 Order/Create