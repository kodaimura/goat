# GOAT
## Go Web Application Template
Go(Gin)のWebアプリケーション雛形。詳しくはプログラム参照。  
accountテーブルをデフォルトで用意しており、サインアップ/ログイン機能を実装。

## Install
```
$ git clone https://github.com/kodaimura/goat
```
* goat/bin にPATHを通す
* 実行権限付与

```
$ chmod -R +x path/to/goat/bin
```

## Usage
### プロジェクト作成
```
$ goat-create-app <appname> [-db {sqlite3| postgres | mysql}]
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
