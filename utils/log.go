/*
 * @Description:
 * @Author: maxyang
 * @Date: 2022-01-06 13:48:17
 * @LastEditTime: 2022-01-10 13:43:42
 * @LastEditors: liutq
 * @Reference:
 */
package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/astaxie/beego/logs"
)

func WriteLog(info string, level string) {
	_dir := "./log"
	exist, err := PathExists(_dir)
	if err != nil {
		fmt.Println("get dir error!", err)
		return
	}

	//如果文件夹不存在，创建文件夹
	if !exist {
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			fmt.Println("make dir error!", err)
			return
		}
	}

	config := make(map[string]interface{})
	config["filename"] = fmt.Sprintf("./log/%s.log", ConfObj.LogFileName)
	config["level"] = logs.LevelDebug

	configStr, err := json.Marshal(config)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))

	if level == "W" {
		logs.Warn("[coinTools] %s", info)
	} else if level == "E" {
		logs.Error("[coinTools] %s", info)
	} else {
		logs.Trace("[coinTools] %s", info)
	}
	return
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
