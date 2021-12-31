package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "admin:admin@tcp(192.168.168.138:3306)/gozero"

	db, err = sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("open %s failed,err: %v \n", dsn, err)
		return err
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed, err: %v \n", dsn, err)
		return err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	fmt.Printf("连接数据库库成功\n")
	return
}

type book struct {
	book  string
	price int
}

func queryRowDemo() {
	sqlStr := "select book, price from book where book=?"
	var b book
	err := db.QueryRow(sqlStr, 1).Scan(&b.book, &b.price)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("book:%s price:%s \n", b.price, book{})
}

func query(price int) {
	sqlStr := "select book, price from book where price > ?"
	rows, err := db.Query(sqlStr, price)
	if err != nil {
		fmt.Printf("query failed,err: %v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var b book
		err := rows.Scan(&b.book, &b.price)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		//fmt.Printf("book:%s price:%v\n", b.book, b.price)
	}
}

func insertRowDemo(book string, price int) {
	sqlStr := "insert into book(book,price) value (?,?)"
	ret, err := db.Exec(sqlStr, book, price)
	if err != nil {
		fmt.Printf("insert failed,err: %v\n", err)
		return
	}
	_, err = ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert book failed,err:%v\n", err)
		return
	}
	//fmt.Printf("insert success,the book is %v.\n", theBook)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("数据库初始化失败\n")
	}
	num := 1000
	offset := 100000
	//start := time.Now()
	for i := num; i < num+offset; i++ {
		//bookName := "bookName" + strconv.Itoa(i)
		//insertRowDemo(bookName, i)
	}
	//fmt.Printf("insert data using time is :%v\n", time.Now().Sub(start))
	fmt.Printf("now time is: %v\n", time.Now())
	time1 := time.Now()
	query(num)
	fmt.Printf("query data using time is :%v\n", time.Now().Sub(time1))
}
