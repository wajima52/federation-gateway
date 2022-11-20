## Subgraph (Go言語)
#### マイグレーション
- 参考: [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- create
```bash
/go/bin/migrate create -ext sql -dir db/migrations -tz "Asia/Tokyo" create_XXX_table
```
- up
```bash
/go/bin/migrate -database ${POSTGRESQL_URL} -path db/migrations up
```

#### Model作成
- 参考: [volatiletech/sqlboiler](https://github.com/volatiletech/sqlboiler)
- モデル作成
```bash
/go/bin/sqlboiler --struct-tag-casing camel psql
```

#### GraphQL
- 参考: [99designs/gqlgen](https://github.com/99designs/gqlgen)
- GraphQL作成手順
1. sqlboilerでモデルを作成
2. gqlgen.ymlの`models`に作成したモデルを追加
3. `schema.graphqls`にクエリを追加
4. `resolver.go`のResolverに問い合わせる関数を追加
5. 以下のコマンドを実行
```bash
go run github.com/99designs/gqlgen generate
```
6. schema.resolvers.goに実装を記述


## Supergraph Schema作成
#### 準備
roverをインストール
```bash
curl -sSL https://rover.apollo.dev/nix/latest | sh
```

#### Supergraph Schema作成
1. 以下のコマンドを実行
```bash
rover supergraph compose --config ./supergraph.yaml > ./router/supergraph.graphql 
```

