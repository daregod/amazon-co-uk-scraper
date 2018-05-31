# amazon.co.uk scraper test task

# Exersise:
>## Тестовая задача для Senior Golang разработчика
>
>
>### Задача
>
>Можно использовать любые библиотеки.
>Результат необходимо залить в любой репозиторий.
>
>- Написать сервис на Golang, который принимает массив URL-ов товаров https://www.amazon.co.uk,
>для данных URL он должен загрузить наименование, цена, фото товара(только URL, грузить изображение не надо), признак наличия, пример:
>
>``` json
>[
>  {
>    "url": "https://www.amazon.co.uk/gp/product/1509836071",
>    "meta": {
>        "title": "The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts",
>        "price": "8.49",
>        "image": "https://images-na.ssl-images-amazon.com/images/I/51kB2nKZ47L._SX382_BO1,204,203,200_.jpg",
>    }
>  },
>  // ...
>]
>```
>- Сервис необходимо завернуть в Docker
>- [Не обязательно] Плюсом будет реализация асинхронной загрузки, когда по некоторому requestID мы можем получить результат

# Install:
1. `go get github.com/daregod/amazon-co-uk-scraper`  
2. `go get github.com/go-task/task/cmd/task`  
3. `task update-deps`

# Run locally:
`task run`

# Run dockerized:
`task docker`

# How to use:
See [API-DESC.md](./API-DESC.md)

# Tests:
## Quick test:
`task test`

## Throttle test with integration:
`task deep-test`
* NOTE1: integration test needs internet connection
* NOTE2: deep tests may need `race-detector-runtime` package

