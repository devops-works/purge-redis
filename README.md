# purge-redis

This tool will purge a specified key on a redis server using `scan`.

## Build

- shell :

```bash
go build -o purge-redis main.go
```

- docker :

```bash
docker build -t devopsworks/purge-redis .
```

## Run

- shell :

```bash
./purge-redis -k 'foo*' -s ${REDIS_HOST}:6379
```

- docker :

```bash
docker run --rm devopsworks/purge-redis -k 'foo*' -s ${REDIS_HOST}:6379 -t 2
```

## Options

- `k` : key to remove (can contain wildcard; watch out for shell expansion !)
- `s` : server, `ip:port` format
- `t` : time, number of minutes to wait for next trigger

## Caveats

Watch out for shell expansion on your `-k` argument !

## Missing

- no DB selection
- no authentication
