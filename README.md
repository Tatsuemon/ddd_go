# DDD入門

- オニオンアーキテクチャ

#### Domain
1. Value Object
2. Entity
3. Service
4. Repository

#### db
- `/schmas/` 以下にcreate文を書いたsqlを配置
```
$ sql-migrate status # status確認
$ sql-migrate up # migrate
$ sql-migrate down # rollback
$ sql-migrate new ????? # migration fileの作成
```

## TODO

- docker-compose
- transaction
- env
