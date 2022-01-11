/*
 * @Description:
 * @Author: maxyang
 * @Date: 2022-01-10 13:41:30
 * @LastEditTime: 2022-01-10 16:15:55
 * @LastEditors: liutq
 * @Reference:
 */
/*
 * @Description:批量生成邮箱
 * @Author: maxyang
 * @Date: 2022-01-10 11:16:42
 * @LastEditTime: 2022-01-10 13:40:18
 * @LastEditors: liutq
 * @Reference:
 */
package core

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/maxyang107/collectcoin/utils"
)

func BatchCreateEmail(cnum int, fileNmae string) {
	fileName := "./static/name_list.txt"        // txt文件路径
	data, err_read := ioutil.ReadFile(fileName) // 读取文件
	if err_read != nil {
		utils.WriteLog("文件读取失败！"+err_read.Error(), "E")
	}
	dataLine := strings.Split(string(data), "\n")
	var emailSuffix = []string{"@163.com", "@gmail.com", "@yahoo.com", "@msn.com", "@hotmail.com", "@aol.com", "@live.com", "@qq.com", "@163.net", "@googlemail.com", "@mail.com", "@walla.com", "@inbox.com", "@ctimail.com"}
	var num = 1
	rand.Seed(time.Now().UnixNano())
	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	for {
		randName := dataLine[rand.Intn(len(dataLine))]
		prex := emailSuffix[rand.Intn(len(emailSuffix))]
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", num), randName+"***"+prex)
		f.SetActiveSheet(index)
		num++
		if num > cnum {
			break
		}
	}
	if err := f.SaveAs(fmt.Sprintf("%s.xlsx", fileNmae)); err != nil {
		utils.WriteLog("保存email文件失败："+err.Error(), "E")
	}

	utils.WriteLog(fmt.Sprintf("批量生成email成功,共计%d", cnum), "T")
}
