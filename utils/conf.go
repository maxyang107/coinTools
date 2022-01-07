/*
 * @Description:配置文件
 * @Author: maxyang
 * @Date: 2022-01-07 17:00:43
 * @LastEditTime: 2022-01-07 17:06:43
 * @LastEditors: liutq
 * @Reference:
 */

package utils

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ChanConn string

	GasLimmit uint64
}

var ConfObj *Config

func init() {
	ConfObj = &Config{
		ChanConn:  "https://bsc-dataseed1.binance.org",
		GasLimmit: 210000,
	}
	data, err := ioutil.ReadFile("conf.json")

	if err == nil {
		json.Unmarshal(data, &ConfObj)
	}
}
