# setting up
```shell
docker-compose up
migrate -database='mysql://user:password@tcp(0.0.0.0:3306)/real_world' -path=migration up 1
```
