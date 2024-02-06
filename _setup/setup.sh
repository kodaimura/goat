#!/bin/bash

if [ $# -eq 0 ]; then
    echo "ERROR: bash _setup/setup.sh <appname> [-db {sqlite3| postgres | mysql}]"
    exit 1
fi

APP_NAME=${1:-goatapp}
shift

SELF_DIR=$(cd $(dirname $0); pwd)
DB="sqlite3"

while (( $# > 0 ))
do
  case $1 in
    -db | --db | --db=*)
      if [[ $1 =~ ^--db= ]]; then
        DB=$2
      elif [ -z $2 ]; then
        echo "ERROR: bash _setup/setup.sh <appname> [-db {sqlite3| postgres | mysql}]"
        exit 1
      else
        DB=$2
        shift
      fi
  esac
  shift
done

if [ $DB != "sqlite3" ] && [ $DB != "postgres" ] && [ $DB != "mysql" ]; then
	echo "ERROR: bash _setup/setup.sh <appname> [-db {sqlite3| postgres | mysql}]"
  exit 1
fi

cd $SELF_DIR
cd ../

mv cmd/goat cmd/$APP_NAME

mkdir log

if [ $DB = "sqlite3" ]; then
  touch $APP_NAME.db
  sqlite3 $APP_NAME.db < ./scripts/create-table.sql
fi

if [ $DB = "postgres" ]; then
  rm -r scripts
  cp -r _setup/postgresql/scripts .
  cp -r _setup/postgresql/repository internal/
  cp -r _setup/postgresql/db internal/core/
  cp -r _setup/postgresql/env config/
  cp -r _setup/postgresql/docker-compose.yml .
  cp -r _setup/postgresql/Dockerfile .
fi

if [ $DB = "mysql" ]; then
  rm -r scripts
  cp -r _setup/mysql/scripts .
  cp -r _setup/mysql/repository internal/
  cp -r _setup/mysql/db internal/core/
  cp -r _setup/mysql/env config/
  cp -r _setup/mysql/docker-compose.yml .
  cp -r _setup/mysql/Dockerfile .
  cp -r _setup/mysql/my.ini .
fi

# goatを置換
for fpath in `find . -name "*.go"`
do sed -i "" s/goat/$APP_NAME/g $fpath
done
sed -i "" s/goat/$APP_NAME/g go.mod
sed -i "" s/goat/$APP_NAME/g Makefile
sed -i "" s/goat/$APP_NAME/g Dockerfile
sed -i "" s/goat/$APP_NAME/g docker-compose.yml
sed -i "" s/goat/$APP_NAME/g ./config/env/local.env
sed -i "" s/goat/$APP_NAME/g ./config/env/docker.env
sed -i "" s/goat/$APP_NAME/g ./web/static/manifest.json

for fpath in `find . -name "*.DS_Store"`
do rm $fpath
done

cat <<EOF > .gitignore
*.log
*.db
*.sqlite3
*.env
.DS_Store
main
data
EOF

yes | rm -r _setup