/*
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Date: 2020-11-09 17:31:23
 * @LastEditTime: 2020-11-11 15:34:22
 * @LastEditors: ferried
 * @Description: Basic description
 * @FilePath: /bidding-go/nmgggzyjy/table.go
 * @LICENSE: Apache-2.0
 */
package nmgggzyjy

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/ferried/bidder/util"
	"github.com/ferried/bidder/wechat"
)

var (
	// TableSQL 创建表语句
	TableSQL string = "CREATE TABLE `nmgggzyjy` (`GUID` char(255) NOT NULL,`Number` char(255) DEFAULT NULL,`Title` char(255) DEFAULT NULL,`Date` char(255) DEFAULT NULL,`Link` char(255) DEFAULT NULL,PRIMARY KEY (`GUID`)) ENGINE=InnoDB DEFAULT CHARSET=utf8;"
)

// CreateTable 创建表
func CreateTable(db *sql.DB) {
	util.Println("正在创建表:" + TableSQL)
	db.Exec(TableSQL)
}

// SaveEntity 保存实体到库中
func SaveEntity(db *sql.DB, data []Entity) {
	if ckdata := CheckEntity(db, data); len(ckdata) > 0 {
		SQLS := make([]string, 0)
		for _, v := range ckdata {
			tempSQL := fmt.Sprintf("('%s','%s','%s','%s','%s')", v.GUID, v.Number, v.Title, v.Date, v.Link)
			SQLS = append(SQLS, tempSQL)
		}
		sql := fmt.Sprintf("INSERT INTO nmgggzyjy(GUID,Number,Title,Date,Link) values %s", strings.Join(SQLS, ","))

		if _, err := db.Exec(sql); err != nil {
			panic(err)
		}
		util.Println(fmt.Sprintf("记录已入库:%v", ckdata))

		entitiesPush := []map[string]interface{}{}
		for _, v := range ckdata {
			if strings.Contains(v.Title, "学校") || strings.Contains(v.Title, "学院") || strings.Contains(v.Title, "大学") || strings.Contains(v.Title, "院校") {
				entityPush := map[string]interface{}{}
				entityPush["LINK"] = PagePrefix + v.Link
				entityPush["TITLE"] = v.Title
				entityPush["TITLENEW"] = v.Title
				entityPush["INFOD"] = v.Number
				entityPush["INFODATE"] = v.Date
				entityPush["CONTENT"] = "请点击查看详情..."
				entityPush["CATEGORYNAME"] = "暂无"
				entityPush["CATENAME"] = "暂无"
				entitiesPush = append(entitiesPush, entityPush)
			}
		}
		wechat.Client.Push(entitiesPush)
	}

}

// CheckIsBlankTable 查看表是否没数据
func CheckIsBlankTable(db *sql.DB) bool {
	util.Println("正在检查表中是否有初始化数据")
	isBlank := true
	sql := "SELECT * FROM nmgggzyjy"
	r, e := db.Query(sql)
	if e != nil {
		panic(e)
	}
	if r.Next() {
		isBlank = false
	}
	return isBlank
}

// CheckEntity 检查去重
func CheckEntity(db *sql.DB, data []Entity) []Entity {
	util.Println("正在对比数据库中差异化记录")
	// g sql
	sql := "SELECT * FROM nmgggzyjy n where n.GUID in ("
	ids := make([]string, 0)
	for _, val := range data {
		ids = append(ids, val.GUID)
	}
	sql += fmt.Sprintf("'%s')", strings.Join(ids, "','"))
	// execute
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	if rows == nil {
		return data
	}
	// get data
	DBData := make([]Entity, 0)
	for rows.Next() {
		e := Entity{}
		if err := rows.Scan(&e.GUID, &e.Number, &e.Title, &e.Date, &e.Link); err != nil {
			panic(err)
		}
		DBData = append(DBData, e)
	}
	res := make([]Entity, 0)
	for _, v := range data {
		if !Contain(DBData, v) {
			res = append(res, v)
		}
	}
	// cut data to find need to saved data
	util.Println(fmt.Sprintf("差异化记录对比完毕,结果:%v", res))
	return res
}

// Contain 判断包含
func Contain(s1 []Entity, t Entity) bool {
	for _, v := range s1 {
		if v.GUID == t.GUID {
			return true
		}
	}
	return false
}
