### RESTAPI using Golang and Gin

---

1. Get All Albums

Input in Git Bash:

```bash
$ curl localhost:8080/albums
```

Output:

```json
[
  {
    "id": "1",
    "title": "Abbey Road: Super Deluxe Edition",
    "artist": "The Beatles",
    "price": 100.5
  },
  {
    "id": "2",
    "title": "The Car",
    "artist": "Arctic Monkeys",
    "price": 39.99
  },
  {
    "id": "3",
    "title": "Is This It",
    "artist": "The Strokes",
    "price": 49.99
  }
]
```

2. Get Specific Album

URL: localhost:8080/albums/:id

Input in Git Bash:

```bash
$ curl localhost:8080/3
```

Output:

```json
{
  "id": "2",
  "title": "The Car",
  "artist": "Arctic Monkeys",
  "price": 39.99
}
```

3. Create

URL: localhost:8080/albums

Input in Git Bash:

```bash
$ curl http://localhost:8080/albums \
> --include \
> --header "Content-Type: application/json" \
> --request "POST" \
> --data '{ "id": "4", "title": "Unlimited Love", "artist": "Red Hot Chili Peppers", "price": 79.99 }'
```

Output in Get All Albums:

```json
[
  {
    "id": "1",
    "title": "Abbey Road: Super Deluxe Edition",
    "artist": "The Beatles",
    "price": 100.5
  },
  {
    "id": "2",
    "title": "The Car",
    "artist": "Arctic Monkeys",
    "price": 39.99
  },
  {
    "id": "3",
    "title": "Is This It",
    "artist": "The Strokes",
    "price": 49.99
  },
  {
    "id": "4",
    "title": "Unlimited Love",
    "artist": "Red Hot Chili Peppers",
    "price": 79.99
  }
]
```

4. Update

URL: localhost:8080/albums/[:id]

Input in Git Bash:

```bash
$ curl http://localhost:8080/albums/4 \
> --include \
> --header "Content-Type: application/json" \
> --request "PUT" \
> --data '{ "price": 49.99 }'
```

Output:

```json
{
  "id": "4",
  "title": "Unlimited Love",
  "artist": "Red Hot Chili Peppers",
  "price": 49.99
}
```

5. Delete

URL: localhost:8080/albums/[:id]

Input in Git Bash:

```bash
$ curl http://localhost:8080/albums/4 \
> --include \
> --request "DELETE"
```

Output in Git Bash:

```bash
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 30 Mar 2024 02:50:40 GMT
Content-Length: 27

{"message":"album deleted"}
```

Output in Get All Albums:

```json
[
  {
    "id": "1",
    "title": "Abbey Road: Super Deluxe Edition",
    "artist": "The Beatles",
    "price": 100.5
  },
  {
    "id": "2",
    "title": "The Car",
    "artist": "Arctic Monkeys",
    "price": 39.99
  },
  {
    "id": "3",
    "title": "Is This It",
    "artist": "The Strokes",
    "price": 49.99
  }
]
```
