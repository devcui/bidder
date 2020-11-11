/*
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Date: 2020-11-09 10:05:42
 * @LastEditTime: 2020-11-10 17:58:24
 * @LastEditors: ferried
 * @Description: Basic description
 * @FilePath: /bidding-go/nmgggzyjy/parse.go
 * @LICENSE: Apache-2.0
 */
package nmgggzyjy

import (
	"fmt"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/ferried/bidder/util"
	"golang.org/x/net/html"
)

// HandleRow 格式化node获取值
func HandleRow(node *html.Node) Entity {
	util.Println("正在分析当前页数抓取结果")
	numberNode, err := htmlquery.Query(node, "/td[2]")
	handleError(err)
	number := htmlquery.CreateXPathNavigator(numberNode).Value()

	titleNode, err := htmlquery.Query(node, "td[3]/a")
	handleError(err)
	title := htmlquery.CreateXPathNavigator(titleNode).Value()

	hrefNode, err := htmlquery.Query(node, "td[3]/a/@href")
	handleError(err)
	href := htmlquery.CreateXPathNavigator(hrefNode).Value()

	publishDateNode, err := htmlquery.Query(node, "/td[4]")
	handleError(err)
	publishDate := htmlquery.CreateXPathNavigator(publishDateNode).Value()

	guid := strings.Split(href, "guid=")[1]
	util.Println(fmt.Sprintf("分析抓取结果完毕 %s %s %s %s %s", guid, number, title, href, publishDate))
	return Entity{
		GUID:   guid,
		Number: number,
		Title:  title,
		Link:   href,
		Date:   publishDate,
	}
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
