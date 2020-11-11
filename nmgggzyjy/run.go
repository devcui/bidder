/*
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Date: 2020-11-06 13:44:28
 * @LastEditTime: 2020-11-10 17:54:24
 * @LastEditors: ferried
 * @Description: Basic description
 * @FilePath: /bidding-go/nmgggzyjy/run.go
 * @LICENSE: Apache-2.0
 */
package nmgggzyjy

import (
	"database/sql"

	"github.com/antchfx/htmlquery"
	"github.com/ferried/bidder/util"
)

// Run 跑内蒙古公共资源爬虫
func Run() {
	// 如果没表
	if !util.CheckTable("nmgggzyjy") {
		// 自动建表
		util.Execute(CreateTable)
	}
	// 如果表里没数据那么添加10页数据
	util.Execute(func(db *sql.DB) {
		if CheckIsBlankTable(db) {
			pullData := Pull(InitPages)
			SaveEntity(db, pullData)
		}
	})
	//正常拉取数据
	pullData := Pull(PullPages)
	util.Execute(func(db *sql.DB) {
		SaveEntity(db, pullData)
	})
}

// Pull 拉取数据
func Pull(p int) []Entity {
	// 需要对比的结果集
	dataset := make([]Entity, 0)
	// 每次抓3页(0-1-2)
	for page := 1; page < p; page++ {
		// 抓取table下所有tr,tr从第二行开始
		cNode := PageXpathNode(page, "/html/body/div[2]/div[2]/div/div[4]/table/tbody/tr[2]")
		for {
			if cNode.Current() != nil && cNode.Current().Data == "tr" {
				dataset = append(dataset, HandleRow(cNode.Current()))
				cNode = htmlquery.CreateXPathNavigator(cNode.Current().NextSibling)
			} else {
				break
			}
		}
	}
	return dataset
}
