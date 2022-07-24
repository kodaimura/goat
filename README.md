# GOAT
## Go Web Application Template
Go(Gin)のWebアプリケーション雛形作成スクリプト。\
ディレクトリ構成 + Signup/Login/Logout 機能を画面およびサーバプログラム自動生成。

* インストール後 ~ /goat/bin にPATHを通し、下記コマンド実行
```
goat <appname> [-db sqlite3 or pg]
```

* オプションについては下記にも対応
* オプションを省略した場合は sqlite3 が選択される
```
 --db sqlite3 [pg]
 --db=sqlite3 [pg]
```

## 初期設定
### appname/config/env 内のファイルを修正

```
# local.env (開発環境用の設定ファイル)

APP_HOST=localhost (必須)
APP_PORT=8080      (必須)
DB_NAME=           (必須)
DB_HOST=localhost  (pg の場合必須)
DB_PORT=5432       (pg の場合必須)
DB_USER=           (pg の場合必須)
DB_PASSWORD=       (pg の場合必須)
JWT_SECRET_KEY=    (必須)
```

```
# .env (本番環境用の設定ファイル)

APP_HOST=          (必須) 
APP_PORT=          (必須)
DB_NAME=           (必須)
DB_HOST=           (pg の場合必須)
DB_PORT=           (pg の場合必須)
DB_USER=           (pg の場合必須)
DB_PASSWORD=       (pg の場合必須)
JWT_SECRET_KEY=    (必須)
```

### DB作成
* sqlite3

```
appname/ 配下にファイル作成
> touch [DB_NAME(←local.env)].db

> sqlite3 DB_NAME.db

DB_NAME> scripts/create-table.sql を実行
```

* postgresql
```
> psql -d postgres

postgres=# CREATE DATABASE [DB_NAME(←local.env)]
CREATE DATABASE

postgres=# \c [DB_NAME]

DB_NAME=# scripts/pg-create-function.sql を実行
DB_NAME=# scripts/pg-create-table.sql を実行
```

### 実行
* 開発環境では下記コマンドで実行 (local.envが読み込まれる)

```
ENV=local go run cmd/myte/main.go
```
