# GOAT
## Go Web Application Template
Go(Gin)のWebアプリケーション雛形。詳しくはプログラム参照。  
accountテーブルをデフォルトで用意しており、サインアップ/ログイン機能を実装。

## Install
### clone
```
$ git clone https://github.com/kodaimura/goat
```
### goat/bin にPATHを通す
```
export PATH=$PATH:path/to/goat/bin
```
### 実行権限付与
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
* goatセットアップ (Dockerコンテナ内)
```
$ make init （初回のみ）
```
* アプリ起動（Dockerコンテナ内）
```
$ make run
```
http://localhost:3000

## Tools
### gent
* model/repository コード自動生成ツール
* make init により、gent 実行ファイルが作成される
* 第一引数にDDLファイルパスを指定して実行する
```
$ ./gent <path/to/create-table.sql>
```
* 第二引数以降にテーブル名を入力し、コードを生成するテーブルを指定可能
```
$ ./gent <path/to/create-table.sql> table1 table2 table3 ...
```
