# Simple Todo API With Go, Gin, Clean Architecture

```bash
go run main.go
```

## Routing in API

Testing in localhost

- Get all todos

```
GET: localhost:8080/items
```

- Get todo item by id

```
GET: localhost:8080/items/:id
```

```
example
localhost:8080/items/1
```

- Create todo item

```
POSTL: localhost:8080/items
```

Body

```
{
  "content": "Testing",
}
```

- Update todo item by id

```
PUT: localhost:8080/items/:id
```

Body

```
{
  "content": "Edited data",
}
```

- Delete todo item by id

```
DELETE: localhost:8080/items/:id
```
