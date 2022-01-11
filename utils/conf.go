/*
 * @Description:配置文件
 * @Author: maxyang
 * @Date: 2022-01-07 17:00:43
 * @LastEditTime: 2022-01-10 16:25:42
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

	LogFileName string
}

var ConfObj *Config

func init() {
	ConfObj = &Config{
		ChanConn:    "https://data-seed-prebsc-1-s1.binance.org:8545",
		GasLimmit:   300000,
		LogFileName: "record",
	}
	data, err := ioutil.ReadFile("./static/conf.json")

	if err == nil {
		json.Unmarshal(data, &ConfObj)
	}
}
