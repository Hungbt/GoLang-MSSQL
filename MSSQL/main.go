package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type User struct {
	name string `json:"name"`
}

func main() {
	db, err := sql.Open("mssql", "server=localhost;port=1433;Database=testdb;User Id=sa;Password=hung")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// GET database
	results, err := db.Query("select * from users")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User

		err = results.Scan(&user.name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.name)
	}

	// --> insert database
	// insert, err := db.Query("INSERT INTO sys.users VALUES ('BuiTuanHung')")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
}
