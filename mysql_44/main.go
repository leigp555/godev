package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	err := initDb()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	fmt.Println("连接数据库成功")
}
func initDb() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/testsql?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	Id       int
	Username string
	Password string
}

//增
func insert(username string, password string) {
	sqlStr := "insert into user_tb1(username,password) values (?,?)"
	ret, err := db.Exec(sqlStr, username, password)
	if err != nil {
		fmt.Printf("insert err: %v\n", err)
	}
	lastInsertId, err2 := ret.LastInsertId()
	if err2 != nil {
		fmt.Printf("get lastInsertId err2: %v\n", err2)
	}
	fmt.Printf("theId: %v\n", lastInsertId)
}

//查
func searchRow() {
	type User struct {
		Id       int
		Username string
		Password string
	}
	var user User
	sqlStr := "select * from user_tb1 where id =?"
	err := db.QueryRow(sqlStr, 10).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		fmt.Printf("查询失败%v\n", err)
	}
	fmt.Println(user)
}
func searchMore() {
	type User struct {
		Id       int
		Username string
		Password string
	}
	var u User
	sqlStr := "select * from user_tb1 where id <>?"
	row, err := db.Query(sqlStr, 5)
	defer row.Close()
	if err != nil {
		fmt.Printf("查询失败%v\n", err)
		return
	}
	for row.Next() {
		row.Scan(&u.Id, &u.Username, &u.Password)
		fmt.Println(u)
	}
}

//改
func update() {
	sqlStr := "update user_tb1 set username=?, password=? where id=?"
	ret, err := db.Exec(sqlStr, "王麻子", "77777", 10)
	if err != nil {
		fmt.Printf("更新失败%v\n", err)
	}
	lastIId, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("id获取失败%v\n", err)
	}
	fmt.Println(lastIId)
}

//删
func delete()  {
	sqlStr := "delete from user_tb1 where username=?"
	ret, err := db.Exec(sqlStr, "tom")
	if err != nil {
		fmt.Printf("删除失败%v\n", err)
	}
	effect, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("影响%v行\n", err)
	}
	fmt.Println(effect)
}
func main() {
 insert("qq","4567")
}
