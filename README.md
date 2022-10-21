# Go RestAPI With Docker

## API 一覧

| 機能         | メソッド | URI     | 権限 |
| ------------ | -------- | ------- | ---- |
| 文字列を返す | GET      | api/v1/ | なし |

done!

| 機能     | メソッド | URI           | 権限 |
| -------- | -------- | ------------- | ---- |
| 新規登録 | POST     | api/v1/signup | なし |
| ログイン | POST     | api/v1/signin | なし |

| 機能               | メソッド | URI               | 権限 |
| ------------------ | -------- | ----------------- | ---- |
| 習慣を登録する     | POST     | api/v1/create     | 有り |
| 習慣を削除する     | DELETE   | api/v1/delete/:id | 有り |
| 習慣を更新する     | POST     | api/v1/update     | 有り |
| 習慣を全て取得する | POST     | api/v1/get        | 有り |

## 使用する流れ

1, まずルートディレクトリに.env ファイルを用意する。

```shell
    touch .env
```

2, ローカルから MySQL のコンテナを作成後、コンテナの MySQL を操作する

```shell
    docker-compose up -d # コンテナの作成 -d -> バックグランドで実行
    docker exec -it mysql_db bash # コンテナに入る
    mysql -u root -p # DBに接続
    show databases; # DBを表示
    use test; # 使用するDBを選択
    show columns from [テーブル名]; # 使用するテーブル名を表示
```

## 仕様

- Push、PL 後 Github Actions にて Lint チェックを行う

## 使用パッケージなど

環境変数の管理

- `https://github.com/joho/godotenv`

ORM

- `https://github.com/go-gorm/gorm`

MySQL

- `https://github.com/go-sql-driver/mysql`

Validation

- `https://github.com/go-playground/validator`

認証関連

- `https://github.com/golang-jwt/jwt`
