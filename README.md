# DDD入門

- オニオンアーキテクチャ

#### Domain
1. Value Object
2. Entity
3. Service
4. Repository

#### db
```
$ docker-compose run app sql-migrate status # status確認
$ docker-compose run app sql-migrate up # migrate
$ docker-compose run app sql-migrate down # rollback
$ docker-compose run app sql-migrate new ????? # migration fileの作成
```

## TODO

- docker-compose
- transaction
- env
