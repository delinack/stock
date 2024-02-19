# stock

API для работы с товарами на складе

# Makefile

`make up` запуск приложения с помощью docker-контейнеров: поднятие базы, накатывание миграций, заполнение базы данными

`make db` поднятие базы _stock_ локально

`make goose-up` применение миграций

`make goose-down` отмена последней миграции

`make linter` установка и прогон линтера

`make format` установка и применение тулзы для форматирования импортов под линтер (smartimports)

`make bin-deps` установка зависимостей (smartimports, linter, goose, gomock) в директиву bin

`make clean` удаление директивы bin и зависимостей (smartimports, linter, goose, gomock)

`make mocks` установка и генерация моков для базы

`make test` прогон тестов

# curl команды с ответами

1. Резервирование товара на складе для доставки

request:
```
curl -X POST http://localhost:8080 -H 'content-type: application/json' -d
'{
    "jsonrpc": "2.0", 
    "method": "Item.ReserveItemsForDelivery", 
    "params": [
        {
        	"stock_id": "5fe06170-4fb3-429a-b950-1ae1a037376e",
        	"items": [
        	    {
        	        "id": "f2aaa171-0118-497b-a0f3-bcffb82533f1",
        	        "quantity": 2
        	    }
        	]
        }
    ],
    "id": 1
}'
```
response:

```json
{
  "result": "items have been reserved",
  "error": null,
  "id": 1
}
```

2. Освобождение резерва товаров

request:
```
curl -X POST http://localhost:8080 -H 'content-type: application/json' -d
'{
    "jsonrpc": "2.0", 
    "method": "Item.DeleteItemsReservation", 
    "params": [
    	{
        	"stock_id": "5fe06170-4fb3-429a-b950-1ae1a037376e",
        	"items": [
        	    {
        	        "id": "f2aaa171-0118-497b-a0f3-bcffb82533f1",
        	    }
        	]
        }
    ],
    "id": 1
}'
```

response:
```json
{
  "result": "item's reservation have been deleted",
  "error": null,
  "id": 1
}
```

3. Получение количества оставшихся товаров на складе

request:
```
curl -X POST http://localhost:8080 -H 'content-type: application/json' -d
'{
    "jsonrpc": "2.0", 
    "method": "Stock.GetItemsQuantity", 
    "params": [
        {
            "stock_id": "719ece6f-65fc-4cc2-b542-7dc673c6c6a8"
        }
    ], 
    "id": 1
}'
```

response:
```json
{
	"result": [
	  {
		"item_id": "f2aaa171-0118-497b-a0f3-bcffb82533f1",
		"quantity": 5
	  },
	  {
		"item_id": "00ef538f-02c2-4a23-8695-385918026262",
		"quantity": 43
	  },
	  {
		"item_id": "36011ed2-2ae7-4ee9-8567-5900b7b230e3",
		"quantity": 12
	  }
	],
    "error": null,
    "id": 1
}
```