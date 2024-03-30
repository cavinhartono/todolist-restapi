### RESTAPI using Golang

---

1. Get All Albums

Input:

```bash
curl localhost:8080/albums
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

Input:

```bash
curl localhost:8080/3
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

Input:
