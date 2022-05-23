package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

//初始化连接数据库
var db *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/testsql?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Panic("failed to connect database")
	}
	db = d
	fmt.Println("数据库连接成功")
}

//创建表
func createTable() {
	db.AutoMigrate(&Product{})
}

//增加一条记录
func addRecord() {
	db.Create(&Product{Code: "宝马", Price: 200000})
}

//查询记录
func searchRecord() {
	var product Product
	db.First(&product, 1) // 根据整型主键查找   //查询结果放到哪个结构体上，查询条件
	fmt.Printf("主键查询第一条数据%v\n", product)

	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录 //条件查询
	fmt.Printf("主键查询第一条数据%v\n", product)
}

//更新记录
func updateRecord() {
	var product Product
	defer searchRecord()
	// Update - 将 product 的 price 更新为 200
	db.First(&product, 1) //更新之前先把他查出来
	//db.Model(&product).Update("Price", 200)

	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "G63"}) // 仅更新非零值字段 通过结构体更新
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"}) //通过map更新
}

// Delete - 删除 product
func deleteRecord() {
	//删除不是物理删除只是添加了删除标记
	var product Product
	db.First(&product, 2) //删除之前先把他查出来

	db.First(&product, "code = ?", "宝马")
	db.Delete(&product, 1)
}
func main() {
	//createTable()
	//addRecord()
	//searchRecord()
	//updateRecord()
	//deleteRecord()
	type CreditCard struct {
		gorm.Model
		Number string
		UserID uint
	}

	type User struct {
		gorm.Model
		Name       string
		CreditCard CreditCard
	}
	err := db.AutoMigrate(&CreditCard{})
	if err != nil {
		log.Panic(err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Panic(err)
	}
	db.Create(&User{
		Name:       "jinzhu",
		CreditCard: CreditCard{Number: "411111111111"},
	})
}
