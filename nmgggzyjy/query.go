package nmgggzyjy

/*
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Date: 2020-11-05 09:40:12
 * @LastEditTime: 2020-11-10 17:55:09
 * @LastEditors: ferried
 * @Description: Basic description
 * @FilePath: /bidding-go/nmgggzyjy/query.go
 * @LICENSE: Apache-2.0
 */

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/antchfx/htmlquery"
	"github.com/ferried/bidder/util"
)

// PageQuery 查询分页
func PageQuery(page int) []byte {
	util.Println(fmt.Sprintf("正在抓取第%v页内容", page))
	r, err := ioutil.ReadAll(request(PageURL, page))
	if err != nil {
		panic(err)
	}
	return r
}

// PageXpathValue Xpath查询分页内容
func PageXpathValue(page int, path string) string {
	body := PageQuery(page)
	w := bytes.NewBuffer(body)
	r := bytes.NewBuffer(make([]byte, 0))
	util.Minify(r, w)
	node, err := htmlquery.Parse(r)
	if err != nil {
		panic(err)
	}
	qn, err := htmlquery.Query(node, path)
	if err != nil {
		panic(err)
	}
	return htmlquery.CreateXPathNavigator(qn).Value()
}

// PageXpathNode 通过xpath查询Html节点
func PageXpathNode(page int, path string) *htmlquery.NodeNavigator {
	body := PageQuery(page)
	w := bytes.NewBuffer(body)
	r := bytes.NewBuffer(make([]byte, 0))
	util.Minify(r, w)
	document, err := htmlquery.Parse(r)
	if err != nil {
		panic(err)
	}
	node, err := htmlquery.Query(document, path)
	if err != nil {
		panic(err)
	}
	return htmlquery.CreateXPathNavigator(node)
}

// request 发送http请求到内蒙古公共资源交易
func request(uri string, page int) io.ReadCloser {
	pPage := fmt.Sprintf("%d", page)
	scroll := fmt.Sprintf("%d", 1109)
	res, err := http.PostForm(uri, url.Values{"currentPage": {pPage}, "scrollValue": {scroll}})
	if err != nil {
		panic(err)
	}
	return res.Body
}
