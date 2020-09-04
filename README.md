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
$ skeema diff # 差分
$ skeema push # dbに反映
$ skeema pull # dbから反映
```