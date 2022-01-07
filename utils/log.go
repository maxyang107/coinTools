/*
 * @Description:
 * @Author: maxyang
 * @Date: 2022-01-06 13:48:17
 * @LastEditTime: 2022-01-07 14:50:01
 * @LastEditors: liutq
 * @Reference:
 */
package utils

// import (
// 	"log"
// 	"os"
// )

// func Info(info string) {
// 	logfile, err := os.OpenFile("log/op.log", os.O_RDWR|os.O_CREATE, 0666)
// 	if err != nil {
// 		logfile, err = os.Create("log/op.log")
// 		if err != nil {
// 			log.Fatalln(err)
// 			return
// 		}
// 	}
// 	defer logfile.Close()
// 	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
// 	logger.Println(info)
// }

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var Loger *log.Logger

func init() {
	filename := fmt.Sprintf("./%s.log", time.Unix(time.Now().Unix(), 0).Format("20060102150405"))
	filename = strings.TrimSuffix(filename, "\n")
	filename = strings.TrimSuffix(filename, "\r")
	logfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		logfile, err = os.Create(filename)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
	Loger = log.New(logfile, "[coinTools]", log.LstdFlags|log.Lshortfile|log.LUTC) // 将文件设置为loger作为输出
	return
}
