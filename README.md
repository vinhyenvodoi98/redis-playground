# redis-cluster
## What
We have 3 redis server
- redis-0 # port 6379
- redis-1 # port 6380
- redis-2 # port 6381

Redis-0 is master ( can write ) and redis-1 and redis-2 is slayer ( only read )
## Run
```
docker-compose up
```
## Test
```
redis-cli -p xxxx -a xxxx
```
