# Counter

Counter is a web application that increases the value of a counter every time
it is clicked.

## Configuration

### Mount Point

The service can be mounted with a specific mount point (prefix path). This path
can be forwarded from the load balancer as the header `X-Mount-Point`
or it can be given as a environment variable, `COUNTER_MOUNT_POINT`, when
starting the container.

### Stores

The application supports four different stores, they are configured with
environment variables.

#### Memory

This is the default store. It stores the counter in memory and is reset on
every restart.

```
# With binary
$ counter-linux

# With docker
$ docker run -d andersjanmyr/counter
```


#### Redis

Stores the counter in Redis. It is configured by setting the environment
variable `REDIS_URL` to a Redis endpoint such as `localhost:6379`

```
# With binary and local Redis
$ REDIS_URL=localhost:6379 counter-linux

# With docker and local Redis
$ docker run -d -e REDIS_URL=localhost:6379 andersjanmyr/counter

# Or with link, assuming a Redis container named redis is running
docker run -d --link redis -e REDIS_URL=redis:6379 andersjanmyr/counter
```

#### Mongo

Stores the counter in Mongo. It is configured by setting the environment
variable `MONGO_URL` to a Mongo endpoint such as `localhost:27017`

```
# With binary and local Mongo
$ MONGO_URL=localhost:27017 counter-linux

# With docker and local Mongo
docker run -d -e MONGO_URL=localhost:27017 andersjanmyr/counter

# Or with link, assuming a Mongo container named mongo is running
docker run -d --link mongo -e MONGO_URL=mongo:27017 andersjanmyr/counter
```

#### Postgres

Stores the counter in Postgres. It is configured by setting the environment
variable `POSTGRES_URL` to a postgres endpoint such as `postgres://localhost`

```
# With binary and local Postgres
$ POSTGRES_URL=postgres://localhost counter-linux

# With docker and local Postgres
docker run -d -e POSTGRES_URL=postgres://localhost andersjanmyr/counter

# Or with link, assuming a Postgres container named postgres is running
docker run -d --link postgres -e POSTGRES_URL=postgres://postgres@postgres andersjanmyr/counter
```

