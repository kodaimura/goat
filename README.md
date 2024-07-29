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

### Dockerで起動
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
$ go mod tidy （初回のみ）
$ make run
```
http://localhost:3000
