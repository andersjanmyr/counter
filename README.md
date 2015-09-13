# Counter

Counter is a web application that simply increases the value of a counter
every time it is clicked. The application supports four different stores.

## Memory

This is the default store. It stores the counter in memory and is reset on
every restart.

## Redis

Stores the counter in Redis. It is configured by setting the environment
variable `REDIS_URL` to a Redis endpoint such as `localhost:6379`

```
docker run -d -e REDIS_URL=localhost:6379 counter
```

## Mongo

Stores the counter in Mongo. It is configured by setting the environment
variable `MONGO_URL` to a Mongo endpoint such as `localhost:27017`

```
docker run -d -e MONGO_URL=localhost:27017 counter
```

## Postgres

Stores the counter in Postgres. It is configured by setting the environment
variable `POSTGRES_URL` to a postgres endpoint such as `postgres://localhost`

```
docker run -d -e POSTGRES_URL=postgres://localhost counter
```

