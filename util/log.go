/*
 * @Author: ferried
 * @Email: harlancui@outlook.com
 * @Date: 2020-11-10 17:44:22
 * @LastEditTime: 2020-11-11 09:51:20
 * @LastEditors: ferried
 * @Description: Basic description
 * @FilePath: /bidding-go/util/log.go
 * @LICENSE: Apache-2.0
 */
package util

import (
	"fmt"
	"log"
	"os"
)

var (
	logger *log.Logger
)

func init() {
	if logger == nil {
		logger = log.New(os.Stdout, "[default] ", log.LstdFlags)
	}
}

// Println ..
func Println(v ...interface{}) {
	if logger != nil {
		logger.Output(3, fmt.Sprintln(v...))
	}

}
