
curl --location --request POST 'http://localhost:8080/cart'

curl --location --request PUT 'http://localhost:8080/cart?basket_id=4' \
--header 'basket_id: 4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "product":"PEN"
}'

curl --location --request PUT 'http://localhost:8080/cart?basket_id=4' \
--header 'basket_id: 4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "product":"TSHIRT"
}'

curl --location --request PUT 'http://localhost:8080/cart?basket_id=4' \
--header 'basket_id: 4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "product":"PEN"
}'

curl --location --request PUT 'http://localhost:8080/cart?basket_id=4' \
--header 'basket_id: 4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "product":"PEN"
}'

curl --location --request PUT 'http://localhost:8080/cart?basket_id=4' \
--header 'basket_id: 4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "product":"MUG"
}'

curl --location --request PUT 'http://localhost:8080/cart?basket_id=4' \
--header 'basket_id: 4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "product":"TSHIRT"
}'

curl --location --request PUT 'http://localhost:8080/cart?basket_id=4' \
--header 'basket_id: 4' \
--header 'Content-Type: application/json' \
--data-raw '{
    "product":"TSHIRT"
}'

curl --location --request GET 'http://localhost:8080/cart?basket_id=4' \
--header 'basket_id: 4' \
--data-raw ''

curl --location --request DELETE 'http://localhost:8080/cart?basket_id=4' \
--header 'basket_id: 4' \
--data-raw ''