/*
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Date: 2020-11-06 15:15:02
 * @LastEditTime: 2020-11-10 17:56:47
 * @LastEditors: ferried
 * @Description: Basic description
 * @FilePath: /bidding-go/nmgggzyjy/const.go
 * @LICENSE: Apache-2.0
 */
package nmgggzyjy

// Const
var (
	PageURL   string = "http://ggzyjy.nmg.gov.cn/jyxx/zfcg/cggg"
	PullPages int    = 4
	InitPages int    = 11
)

// Entity 接收结果实体
type Entity struct {
	GUID   string
	Number string
	Title  string
	Date   string
	Link   string
}
