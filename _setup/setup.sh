#!/bin/bash

APP_NAME=${1:-goatapp}
shift

CURRENT_DIR=$(pwd)
DB="sqlite3"

while (( $# > 0 ))
do
  case $1 in
    -db | --db | --db=*)
      if [[ $1 =~ ^--db= ]]; then
        DB=$2
      elif [ -z $2 ]; then
        echo "'db' requires an argument. (sqlite3 or pg or mysql)" 1>&2
        exit 1
      else
        DB=$2
        shift
      fi
  esac
  shift
done

if [ $DB != "sqlite3" ] && [ $DB != "pg" ] && [ $DB != "mysql" ]; then
	echo "'db' requires an argument. (sqlite3 or pg or mysql)" 1>&2
  exit 1
fi

cd $CURRENT_DIR/$APP_NAME

mv cmd/goat-base cmd/$APP_NAME

mkdir log

if [ $DB = "pg" ]; then
  rm -r scripts
  cp -r _sub/postgresql/scripts .
  cp -r _sub/postgresql/repository internal/
  cp -r _sub/postgresql/db internal/core/
  cp -r _sub/postgresql/env/local.env config/env/
fi

if [ $DB = "mysql" ]; then
  rm -r scripts
  cp -r _sub/mysql/scripts .
  cp -r _sub/mysql/repository internal/
  cp -r _sub/mysql/db internal/core/
  cp -r _sub/mysql/env/local.env config/env/
fi

# goat-baseを置換
for fpath in `find . -name "*.go"`
do sed -i "" s/goat-base/$APP_NAME/g $fpath
done
sed -i "" s/goat-base/$APP_NAME/g ./config/env/local.env

for fpath in `find . -name "*.DS_Store"`
do rm $fpath
done

cat <<EOF > .gitignore
*.log
*.db
*.sqlite3
.env
local.env
.DS_Store
main
EOF

rm go.sum
rm go.mod
yes | rm -r _setup