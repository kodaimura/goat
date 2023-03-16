# GOAT
## Go Web Application Template
Go(Gin)のWebアプリケーション雛形作成スクリプト。\
ディレクトリ構成 + Signup/Login/Logout 機能を画面およびサーバプログラム自動生成。

## Install
```
$ git clone https://github.com/kodaimura/goat
```

* goat/bin にPATHを通す。
* 生成用スクリプトに実行権限付与。

```
$ chmod +x [省略]/goat/bin/goat
```

## Usage

```
$ goat <appname> [-db sqlite3| pg | mysql]
```

* オプションを省略した場合は sqlite3 が選択される

## Setting
### appname/config/env 内のファイルを修正

```
# local.env (開発環境用の設定ファイル)

APP_HOST=localhost (必須)
APP_PORT=3000      (必須)
DB_NAME=           (必須)
DB_HOST=localhost
DB_PORT=
DB_USER=
DB_PASSWORD=
JWT_SECRET_KEY=    (必須)
```

### DB作成
* sqlite3

```
appname/ 配下にファイル作成
$ sqlite3 [DB_NAME].db

DB_NAME> .read scripts/create-table.sql
```

* postgresql
```
$ psql -d postgres

postgres=# CREATE DATABASE [DB_NAME]
CREATE DATABASE

postgres=# \c [DB_NAME]

DB_NAME=# scripts/create-function.sql
DB_NAME=# scripts/create-table.sql
```

### 実行
* 開発環境では下記コマンドで実行 (local.envが読み込まれる)

```
ENV=local go run cmd/<appname>/main.go
```
