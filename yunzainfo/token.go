/*
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Date: 2020-11-11 09:35:07
 * @LastEditTime: 2020-11-11 10:55:36
 * @LastEditors: ferried
 * @Description: Basic description
 * @FilePath: /bidding-go/yunzainfo/token.go
 * @LICENSE: Apache-2.0
 */
package yunzainfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/ferried/bidder/util"
)

// YzToken ..
type YzToken struct {
	Token         string
	YunzaiInfoURL string
	GrantType     string
	ClientID      string
	ClientSecret  string
}

var Client *YzToken

func init() {
	if Client == nil {
		Client = &YzToken{
			Token:         "",
			YunzaiInfoURL: os.Getenv("BIDDING_YUNZAINFO_URL"),
			GrantType:     os.Getenv("BIDDING_GRANT_TYPE"),
			ClientID:      os.Getenv("BIDDING_CLIENT_ID"),
			ClientSecret:  os.Getenv("BIDDING_CLIENT_SECRET"),
		}
	}
}

// FlushToken ..
func (yz *YzToken) FlushToken() {
	util.Println("开始刷新云在TOKEN")
	p := url.Values{
		"grant_type":    {yz.GrantType},
		"client_id":     {yz.ClientID},
		"client_secret": {yz.ClientSecret},
	}
	res, err := http.PostForm(fmt.Sprintf("%s/auth/oauth/token", yz.YunzaiInfoURL), p)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	util.Println(fmt.Sprintf("刷新TOKEN完成:%s", string(body)))

	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err == nil {
		Client.Token = dat["access_token"].(string)
	} else {
		panic(err)
	}

}
