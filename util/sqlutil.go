/*
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Date: 2020-11-09 16:23:52
 * @LastEditTime: 2020-11-10 17:56:24
 * @LastEditors: ferried
 * @Description: Basic description
 * @FilePath: /bidding-go/util/sqlutil.go
 * @LICENSE: Apache-2.0
 */
package util

import (
	"database/sql"
	"fmt"
	"os"
)

// DBConfig 生成Mysql配置
func DBConfig() string {
	username := os.Getenv("BIDDING_MYSQL_USERNAME")
	password := os.Getenv("BIDDING_MYSQL_PASSWORD")
	URL := os.Getenv("BIDDING_MYSQL_URL")
	return fmt.Sprintf("%s:%s@tcp(%s)/bidding?charset=utf8", username, password, URL)
}

// Connection 获取一个DB连接
func Connection() *sql.DB {
	db, err := sql.Open("mysql", DBConfig())
	if err != nil {
		panic(err)
	}
	return db
}

// Execute 执行db操作
func Execute(anonymous func(db *sql.DB)) {
	db := Connection()
	anonymous(db)
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

// CheckTable 检查table是否存在
func CheckTable(tableName string) bool {
	Println("检查" + tableName + "是否存在")
	flag := false
	Execute(func(db *sql.DB) {
		result, err := db.Query("SHOW TABLES")
		if err != nil {
			panic(err)
		}
		for result.Next() {
			var Tables_in_bidding string
			if err := result.Scan(&Tables_in_bidding); err != nil {
				panic(err)
			}
			if tableName == Tables_in_bidding {
				flag = true
			}
		}
	})
	return flag
}
