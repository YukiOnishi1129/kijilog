#!/bin/bash

#DB_PASSWORD="pass"
#TEST_DB_ENV_STRING="test"
# コマンド実行時の第2引数以下をARGに格納
#ARG_SECOND=${@:2}
#ARG_THIRD=${@:3}
#ARG_FORTH=${@:4}

case ${1} in
# テーブルからentityを自動生成 (事前にdbコンテナを立ち上げておくこと)
"entity:create")
 echo  === entity create start ===
 sqlboiler psql -c backend/database.toml -o backend/database/entity -p entity --no-tests --wipe
 echo  === entity create end ===
 ;;
"db:migrate")
 echo  === db migrate start ===
 migrate -source file://backend/database/migrations -database 'postgres://user:pass@localhost:5432/KIJILOG_DB?sslmode=disable' up
 echo  === db migrate end ===
 ;;
"db:seed")
 echo  === db seed start ===
 # shellcheck disable=SC2164
 cd backend
 go run database/seed/seed.go
 echo  === db seed start ===
 ;;
"db:rollback")
 echo  === db rollback start ===
 migrate -source file://backend/database/migrations -database 'postgres://user:pass@localhost:5432/KIJILOG_DB?sslmode=disable' down
 echo  === db rollback end ===
 ;;
"db:reset")
 # db:rollback
 echo  === db rollback start ===
 migrate -source file://backend/database/migrations -database 'postgres://user:pass@localhost:5432/KIJILOG_DB?sslmode=disable' down
 echo  === db rollback end ===
 # db:migrate
 echo  === db migrate start ===
 migrate -source file://backend/database/migrations -database 'postgres://user:pass@localhost:5432/KIJILOG_DB?sslmode=disable' up
 echo  === db migrate end ===
 # db:seed
 echo  === db seed start ===
 # shellcheck disable=SC2164
 cd backend
 go run database/seed/seed.go
 echo  === db seed end ===
 ;;
"lint")
 echo  === staticcheck lint start ===
 # shellcheck disable=SC2164
 cd backend
 staticcheck ./...
 echo  === staticcheck lint end ===
 ;;
"test:all")
 echo  === test all start ===
 # shellcheck disable=SC2164
 cd backend
 go test -v ./...
 echo  === test all end ===
 ;;
#"test")
# # shellcheck disable=SC2164
# cd backend
# go test -v "${ARG_SECOND}" "${ARG_THIRD}" "${ARG_FORTH}"
# ;;
"gql")
 echo  === gqlgen generate start ===
 # shellcheck disable=SC2164
 cd backend
 gqlgen generate
 echo  === gqlgen generate end ===
 ;;
esac