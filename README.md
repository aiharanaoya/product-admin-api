# product-admin-api

## about

商品管理 API（Golang 学習用）

## setup

- 以下が導入されていること

| feature | ver    |
| ------- | ------ |
| Go      | 1.16.6 |

- リポジトリのクローン

```sh
$ git clone git@github.com:nao11aihara/product-admin-api.git
$ cd product-admin-api
```

- 起動

```sh
$ go run main.go
```

- ヘルスチェック

```sh
$ curl http://localhost:8080/health_check
```

## DB

- DB 作成

```sh
$ ./shell/createDb.sh
```

- テーブル作成

```sh
$ ./shell/createTable.sh
```

- テーブル削除

```sh
$ ./shell/dropTable.sh
```
