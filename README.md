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
./purge-redis -k 'foo*' -s ${REDIS_HOST}:6379 -t 2 -n #dry mode on
```

- docker :

```bash
docker run --rm devopsworks/purge-redis -k 'foo*' -s ${REDIS_HOST}:6379 -t 2 -n #dry mode on
```

## Options

- `k` : key to remove (default: "") (can contain wildcard; watch out for shell expansion !)
- `s` : server (default: "localhost:6379"), `ip:port` format
- `t` : time (default: 1), number of minutes to wait for next trigger
- `n` : dry run mode (default: false), listing how many to delete

## Caveats

Watch out for shell expansion on your `-k` argument !

## Missing

- no DB selection
- no authentication
