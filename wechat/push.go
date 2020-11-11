/*
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Date: 2020-11-11 11:50:18
 * @LastEditTime: 2020-11-11 15:35:02
 * @LastEditors: ferried
 * @Description: Basic description
 * @FilePath: /bidding-go/wechat/push.go
 * @LICENSE: Apache-2.0
 */

package wechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ferried/bidder/util"
	"github.com/ferried/bidder/yunzainfo"
)

func init() {
	systemKey := os.Getenv("BIDDING_CLIENT_ID")
	eventKey := os.Getenv("BIDDING_EVENT_KEY")
	accounts := strings.Split(os.Getenv("BIDDING_ACCOUNTS"), ",")

	if Client == nil {
		Client = &Wechat{
			SystemKey: systemKey,
			EventKey:  eventKey,
			Accounts:  accounts,
		}
	}

}

var Client *Wechat

// Wechat ..
type Wechat struct {
	SystemKey string
	EventKey  string
	Accounts  []string
}

// Push 推送
func (wechat *Wechat) Push(params []map[string]interface{}) {
	if len(params) >= 0 {
		util.Println("开始发送推送")
		client := &http.Client{}
		pushParam := map[string]interface{}{}
		pushParam["systemKey"] = wechat.SystemKey
		pushParam["fromUserId"] = wechat.Accounts[0]
		pushParam["sendTime"] = time.Now().Unix() * 1000
		pushParam["toUserIds"] = wechat.Accounts
		pushParam["eventKey"] = wechat.EventKey
		for _, ma := range params {
			pushParam["param"] = ma
			jsonParam, err := json.Marshal(pushParam)
			if err != nil {
				println(err)
			}
			requestBody := bytes.NewBuffer(jsonParam)
			request, _ := http.NewRequest("POST", yunzainfo.Client.YunzaiInfoURL+"/message-center-3/messagePush/pushByEvent", requestBody)
			request.Header.Add("Authorization", "bearer "+yunzainfo.Client.Token)
			request.Header.Add("Content-Type", "application/json; charset=utf-8")
			responseBody, err := client.Do(request)
			if err != nil {
				println(err)
			}
			d, err := ioutil.ReadAll(responseBody.Body)
			if err != nil {
				println(err)
			}
			util.Println(fmt.Sprintf("获取推送响应:%s", string(d)))
		}
		util.Println("推送结束")
	}
}
