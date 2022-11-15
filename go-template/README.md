`token` + `redis` , not `JWT` 
- `gin`
- `zap`
- `gorm`
- `casbin`
- `postgres`
- `viper`

# `Docker`
```shell
docker run --name redis -p 6379:6379  --restart=always \
-v D:\Study\data\redis:/data \
-v D:\Study\data\redis\conf\redis.conf:/etc/redis/redis.conf \
-d redis
```

```shell
docker run --name postgres -p 5432:5432  --restart=always \ 
-e POSTGRES_PASSWORD=admin 
-e POSTGRES_USER=noob \
-e POSTGRES_DB=template \
-e PGDATA=/var/lib/postgresql/data/pgdata \
-v D:\Study\data\postgres:/var/lib/postgresql/data \
-d postgres
```

## format code
```shell
gofmt -l -d -w -s .
```