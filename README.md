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
#### supergraph schema作成手順
1. ３つのターミナルそれぞれで以下のコマンドを１つずつ実行
```bash
rover dev -p 5001 --name accounts --schema ./accounts/graph/schema.graphqls --url http://localhost:8082/query
rover dev -p 5001 --name products --schema ./products/graph/schema.graphqls --url http://localhost:8083/query
rover dev -p 5001 --name reviews  --schema ./reviews/graph/schema.graphqls  --url http://localhost:8084/query
```
2. `http://localhost:5001`へアクセスし、`Schema` > `SDL`を開く
3. コピーするかschemaをダウンロード