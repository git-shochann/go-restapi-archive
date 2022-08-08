# Habit Management

RestAPI を Go で試してみる

## API 一覧

---

| 機能         | メソッド | URI     | 権限 |
| ------------ | -------- | ------- | ---- |
| 文字列を返す | GET      | api/v1/ | なし |

| 機能     | メソッド | URI           | 権限 |
| -------- | -------- | ------------- | ---- |
| 新規登録 | POST     | api/v1/signup | なし |
| ログイン | POST     | api/v1/signin | なし |

| 機能               | メソッド | URI           | 権限 |
| ------------------ | -------- | ------------- | ---- |
| 習慣を登録する     | POST     | api/v1/create | 有り |
| 習慣を削除する     | POST     | api/v1/delete | 有り |
| 習慣を更新する     | POST     | api/v1/update | 有り |
| 習慣を全て取得する | POST     | api/v1/get    | 有り |

## 使用する流れ

---

1, まずルートディレクトリに.env ファイルを用意する。

```shell
    touch .env
```

2, ローカルで MySQL のコンテナを作成後、Docker の network を通して、MySQL を操作する

```shell
    docker-compose build
```

## 使用パッケージなど

---

環境変数の管理

- `https://github.com/joho/godotenv`

ORM

- `https://github.com/go-gorm/gorm`

MySQL

- `https://github.com/go-sql-driver/mysql`
