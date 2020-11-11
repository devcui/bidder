package main

import (
	"time"

	"github.com/ferried/bidder/nmgggzyjy"
	"github.com/ferried/bidder/yunzainfo"
	_ "github.com/go-sql-driver/mysql"
)

/*
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Date: 2020-11-05 09:11:49
 * @LastEditTime: 2020-11-11 15:46:56
 * @LastEditors: ferried
 * @Description: Basic description
 * @FilePath: /bidding-go/main.go
 * @LICENSE: Apache-2.0
 */

func main() {
	yunzainfo.Client.FlushToken()
	nmgggzyjy.Run()
	wechatTokenTimmer := time.NewTicker(time.Minute * 23)
	nmgggzyjyTimmer := time.NewTicker(time.Hour)
	for {
		select {
		case <-wechatTokenTimmer.C:
			yunzainfo.Client.FlushToken()
		case <-nmgggzyjyTimmer.C:
			nmgggzyjy.Run()
		}
	}
}
