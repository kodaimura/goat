# GOAT
## Go Web Application Template
Go(Gin)のWebアプリケーション雛形。詳しくはプログラム参照。  
usersテーブルをデフォルトで用意しており、サインアップ/ログイン機能を実装。

https://github.com/kodaimura/create-app の設定により、  
下記 Install ~ 共通セットアップまでが一つのコマンドで可能。

## Install
```
$ git clone https://github.com/kodaimura/goat <appname>
```

## Usage
### 共通セットアップ
```
$ cd <appname>
$ bash _setup/setup.sh <appname> [-db {sqlite3| postgres | mysql}]
```
* -db オプションを省略した場合は sqlite3 が選択される

### Dockerで起動する場合
* Dockerイメージ作成 & コンテナ起動
```
$ make up
```
* Dockerコンテナアクセス
```
$ make in
```
* アプリ起動（Dockerコンテナ内）
```
$ make run
```
http://localhost:3000

### ローカルで起動する場合
* <appname>/config/env/local.env 修正
```
# local.env (開発環境用の設定ファイル)
APP_HOST=localhost
APP_PORT=3000
DB_NAME=
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
JWT_SECRET_KEY=

# sqlite3 の場合 DB_NAME にはdbファイルの絶対パスまたは、プロジェクトのルートフォルダからの相対パスを記載する。
```
* DB作成

* アプリ起動
```
$ go mod tidy （初回のみ）
$ make lrun
```
http://localhost:3000
