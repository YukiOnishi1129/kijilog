package seeders

import (
	"database/sql"
	"fmt"
)

var (
	ins *sql.Stmt
	err error
)

func CreateTestData(con *sql.DB) error {
	// usersテーブルのテストデータ作成
	if err = CreateUserData(con); err != nil {
		fmt.Println("====hh===")
		return err
	}
	// todoテーブルのテストデータ作成
	if err = CreateTodoData(con); err != nil {
		return err
	}
	return nil
}
