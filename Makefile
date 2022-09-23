dynamo:
	docker run -p 8000:8000 -d amazon/dynamodb-local -jar DynamoDBLocal.jar -sharedDb -dbPath .
list-tables:
	aws dynamodb list-tables --endpoint-url http://localhost:8000 --region local
create-table:
	aws dynamodb create-table --cli-input-json file://db/testing.json --region local --endpoint-url http://localhost:8000
add-item:
	aws dynamodb put-item --table-name testing --item file://db/item.json --region local --endpoint-url http://localhost:8000