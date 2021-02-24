package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
)

func main()  {
	db,err := sql.Open("mysql","astaxie:astaxie@/test?charset=utf8")
	fmt.Println(db)
	fmt.Println(err)

}