package main

import (
	"fmt"
	"time"

	"github.com/ferried/bidder/yunzainfo"
	_ "github.com/go-sql-driver/mysql"
)

/*
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Date: 2020-11-05 09:11:49
 * @LastEditTime: 2020-11-11 10:55:56
 * @LastEditors: ferried
 * @Description: Basic description
 * @FilePath: /bidding-go/main.go
 * @LICENSE: Apache-2.0
 */

func main() {
	yunzainfo.Client.FlushToken()
	fmt.Println(yunzainfo.Client.Token)
	// nmgggzyjy.Run()
}

func timmer() {
	wechatTokenTimmer := time.NewTicker(time.Minute * 23)
	nmgggzyjyTimmer := time.NewTicker(time.Hour)
	for {
		select {
		case <-wechatTokenTimmer.C:
			fmt.Println("flush wechat token")
		case <-nmgggzyjyTimmer.C:
			fmt.Println("pull nmgggzyjyTimmer")
		}
	}
}
