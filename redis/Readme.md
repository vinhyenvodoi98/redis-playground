# Dump and restore
Start redis-0

```
docker-compose up -d redis-0
```

Backtup
```
redis-cli

save
```

Get dir
```
CONFIG GET dir
```
