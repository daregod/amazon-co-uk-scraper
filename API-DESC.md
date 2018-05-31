# API
* Default server port is `8007`

---

## Enqueue
#### Request:
`POST /api/enqueue`  
POST Body:  
``` json
[
  "url1",
  "url2",
// ...
  "urlN"
]
```
#### Response:
``` json
{
  "id":"<ID>",
  "status":"PROCEED OK"
}
```
#### Example:
`curl "http://127.0.0.1:8007/api/enqueue" -d "[\"https://www.amazon.co.uk/gp/product/1509836071\",\"http://amazon.co.uk/gp/product/1787125645\"]"`
> ``` json
>{"id":"bd4e6e78-a6d9-4e26-9c3e-b7754fa46147","status":"PROCEED OK"}
>```

---

## Fetch
#### Request:
`GET /api/<ID>`
#### Response:
``` json
{
  "id":"<ID>",
  "status":"<>",
  "error":"<>",
  "data":[{
    "url":"<>",
    "meta":{
      "title":"<>",
      "price":"<>",
      "image":"<>",
      "available":<>
    }
  },
// ...
  ]
}
```  
any field may be optional.
#### Examples:
`curl -s "http://127.0.0.1:8007/api/job/bd4e6e78-a6d9-4e26-9c3e-b7754fa46147"`
> ``` json
>{
>  "id":"bd4e6e78-a6d9-4e26-9c3e-b7754fa46147",
>  "data":[{
>    "url":"https://www.amazon.co.uk/gp/product/1509836071",
>    "meta":{
>      "title":"The Fat-Loss Plan: 100 Quick and Easy Recipes with Workouts",
>      "price":"8.49",
>      "image":"https://images-eu.ssl-images-amazon.com/images/I/71FL39V4zCL._SL110_.jpg",
>      "available":true
>    }
>  },
>  {
>    "url":"https://www.amazon.co.uk/gp/product/1787125645",
>    "meta":{
>      "title":"Go Systems Programming: Master Linux and Unix system level programming with Go",
>      "price":"41.99",
>      "image":"https://images-eu.ssl-images-amazon.com/images/I/41y7-qWywtL._SL110_.jpg",
>      "available":true
>    }
>  }],
>  "status":"OK"
>}
>```

`curl -s "http://127.0.0.1:8007/api/job/59a38b62-7fe1-4e40-9d41-16d72d3e8edc"`
>``` json
>{"id":"59a38b62-7fe1-4e40-9d41-16d72d3e8edc","status":"Error","error":"GET. Job not found: 59a38b62-7fe1-4e40-9d41-16d72d3e8edc"}
>```

---

## Remove
#### Request:
`DELETE /api/deletejob/<ID>`
#### Response:
``` json
{
  "id":"<ID>",
  "status":"<>",
  "error":"<>"
}
```
#### Examples:
`curl -X DELETE http://127.0.0.1:8007/api/deletejob/bd4e6e78-a6d9-4e26-9c3e-b7754fa46147`
>``` json
>{"id":"bd4e6e78-a6d9-4e26-9c3e-b7754fa46147","status":"DELETED OK"}
>```

`curl -X DELETE http://127.0.0.1:8007/api/deletejob/537f59c4-3ba2-40ef-bce5-8489ea1cacc5`
>``` json
>{"id":"537f59c4-3ba2-40ef-bce5-8489ea1cacc5","status":"Error","error":"DELETE. Job not found: 537f59c4-3ba2-40ef-bce5-8489ea1cacc5"}
>```