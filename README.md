# GOAT
## Go Web Application Template
Go(Gin)のWebアプリケーション雛形作成スクリプト。
ディレクトリ構成 + Signup/Login/Logout 機能を画面およびサーバプログラム自動生成。

* インストール後 ~ /goat/bin にPATHを通し、下記コマンド実行。
```
goat <appname> [-db sqlite3 or pg]
```

* オプションについては下記にも対応
```
 --db sqlite3 [pg]
 --db=sqlite3 [pg]
```
* オプションを省略した場合は sqlite3 が選択される

