
# URL Shortener

This allows the input of a website starting with http:// or https:// and returns the user with a short url.
## Installation

Must have docker installed along with docker compose

Clone the project and run docker compose up

```bash
  git clone https://github.com/robertgouveia/URLShortener
  cd URLShortener
  docker-compose up --build
  docker run urlshortener
```

## Tech Stack

GO, Redis, Docker, Tailwind

Docker setup:
```yml
version: '3.8'
services:
  app:
    build: ./src
    ports:
      - "8080:8080"
    depends_on:
      - redis
    environment:
      - REDIS_HOST=redis
      - REDIS_PORT=6379

  redis:
    image: "redis:latest"
    ports:
      - "6379:6379"
    volumes:
      - redis-data:/data

volumes:
  redis-data:
```


Using the `crypto/rand` package:
```go
bytes := make([]byte, 6)
if _, err := rand.Read(bytes); err != nil {
	fmt.Println(err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
	return
}
```

And using `redis` to store the URL's:
```go
shortUrl := "http://localhost:8080/" + base64.URLEncoding.EncodeToString(bytes)

client := redis.GetClient()
ctx := redis.GetContext()

err := client.Set(ctx, shortUrl, url, 0).Err()
if err != nil {
	fmt.Println("set: ", err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
	return
}
```
