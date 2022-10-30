#### マイグレーション
- 参考: [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- create
```aidl
/go/bin/migrate create -ext sql -dir db/migrations -tz "Asia/Tokyo" create_XXX_table
```
- up
```aidl
/go/bin/migrate -database ${POSTGRESQL_URL} -path db/migrations up
```

#### Model作成
- 参考: [volatiletech/sqlboiler](https://github.com/volatiletech/sqlboiler)
- モデル作成
```aidl
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
```aidl
go run github.com/99designs/gqlgen generate
```
6. schema.resolvers.goに実装を記述