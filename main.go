/*
 * @Description:
 * @Author: maxyang
 * @Date: 2022-01-06 10:17:34
 * @LastEditTime: 2022-01-10 18:05:47
 * @LastEditors: liutq
 * @Reference:
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/maxyang107/collectcoin/core"
	"github.com/maxyang107/collectcoin/utils"
)

func main() {
	utils.WriteLog("当前区块链地址："+utils.ConfObj.ChanConn, "T")
LOOP:
	text := menu()
	switch text {
	case "1":
		batchCollectMainCoin()
	case "2":
		batchCollectErc20Coin()
	case "3":
		createWallet()
	case "4":
		airDropErc20()
	case "5":
		airDropNft()
	case "6":
		batchCreateEmailAccount()
	default:
		fmt.Println("无效的操作")
		goto LOOP
	}

}

func menu() string {
	buf := bufio.NewReader(os.Stdin)
	fmt.Println("==============================")
	fmt.Println("=           工具箱           =")
	fmt.Println("==============================")
	fmt.Println("请输入你要执行的任务")
	fmt.Println("")
	fmt.Println("1.主币归集")
	fmt.Println("")
	fmt.Println("2.ERC20代币归集")
	fmt.Println("")
	fmt.Println("3.批量创建钱包")
	fmt.Println("")
	fmt.Println("4.空投ERC20代币")
	fmt.Println("")
	fmt.Println("5.空投NFT")
	fmt.Println("")
	fmt.Println("6.批量生成邮箱")
	fmt.Println("==============================")
	fmt.Println("")

	fmt.Print("请输入选项：")
	fmt.Println("")
	text, err := buf.ReadString('\r')
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(text, "\r", "", -1)
}

//批量创建钱包
func createWallet() {
	bufnum := bufio.NewReader(os.Stdin)
	fmt.Print("请输入您要创建的钱包数量：")
	num, _ := bufnum.ReadString('\r')
	num = strings.Replace(num, "\r", "", -1)
	bufname := bufio.NewReader(os.Stdin)
	fmt.Print("请输入您要文件名称：")
	name, _ := bufname.ReadString('\r')
	name = strings.Replace(name, "\r", "", -1)
	numint, err := strconv.Atoi(num)
	if err != nil || numint <= 0 || name == "" {
		fmt.Println("输入错误", err)
		os.Exit(0)
	}
	core.CreateWallet(numint, name)
}

//主笔批量归集
func batchCollectMainCoin() {
	buffile := bufio.NewReader(os.Stdin)
	fmt.Print("请输入待归集的excel文件名称：")
	filename, _ := buffile.ReadString('\r')
	filename = strings.Replace(filename, "\r", "", -1)

	bufAccount := bufio.NewReader(os.Stdin)
	fmt.Print("请输入收款账户：")
	accountname, _ := bufAccount.ReadString('\r')
	accountname = strings.Replace(accountname, "\r", "", -1)

	if filename == "" || accountname == "" {
		fmt.Println("输入错误")
		os.Exit(0)
	}

	core.CollectCoin(filename, accountname)
}

//批量转erc20代币
func batchCollectErc20Coin() {
	fmt.Println("执行归集erc20代币前，请确认对应钱包账户里面有足够的主币能够支付链上gas费用")
	bufcontract := bufio.NewReader(os.Stdin)
	fmt.Print("请输入erc20代币合约地址：")
	contractename, _ := bufcontract.ReadString('\r')
	contractename = strings.Replace(contractename, "\r", "", -1)

	buffile := bufio.NewReader(os.Stdin)
	fmt.Print("请输入待归集的excel文件名称：")
	filectename, _ := buffile.ReadString('\r')
	filectename = strings.Replace(filectename, "\r", "", -1)

	bufcollectadd := bufio.NewReader(os.Stdin)
	fmt.Print("请输入收款钱包地址：")
	collectadd, _ := bufcollectadd.ReadString('\r')
	collectadd = strings.Replace(collectadd, "\r", "", -1)

	if contractename == "" || filectename == "" || collectadd == "" {
		fmt.Println("输入错误")
		os.Exit(0)
	}

	core.CollectErc20Coin(contractename, filectename, collectadd)
}

//空投ERC20代币
func airDropErc20() {
	fmt.Println("空投erc20代币前，请确认主钱包账户里面有足够的主币能够支付链上gas费用")
	bufcontract := bufio.NewReader(os.Stdin)
	fmt.Print("请输入erc20代币合约地址：")
	contractename, _ := bufcontract.ReadString('\r')
	contractename = strings.Replace(contractename, "\r", "", -1)

	buffile := bufio.NewReader(os.Stdin)
	fmt.Print("请输入空投的excel文件名称：")
	filectename, _ := buffile.ReadString('\r')
	filectename = strings.Replace(filectename, "\r", "", -1)

	buffromadd := bufio.NewReader(os.Stdin)
	fmt.Print("请输入代币转出钱包地址：")
	fromAdd, _ := buffromadd.ReadString('\r')
	fromAdd = strings.Replace(fromAdd, "\r", "", -1)

	bufprive := bufio.NewReader(os.Stdin)
	fmt.Print("请输入转出钱包私钥：")
	privekey, _ := bufprive.ReadString('\r')
	privekey = strings.Replace(privekey, "\r", "", -1)

	if contractename == "" || filectename == "" || fromAdd == "" || privekey == "" {
		fmt.Println("输入错误")
		os.Exit(0)
	}

	core.AirDropErc20Coin(contractename, filectename, fromAdd, privekey)
}

/**
 * @description: 方法描述：批量创建邮箱
 * @Author: maxyang
 * @return {*}
 */
func batchCreateEmailAccount() {
	bufnum := bufio.NewReader(os.Stdin)
	fmt.Print("请输入生成邮箱账号数量：")
	emailNum, _ := bufnum.ReadString('\r')
	emailNum = strings.Replace(emailNum, "\r", "", -1)
	numint, err := strconv.Atoi(emailNum)
	if err != nil {
		utils.WriteLog("输入字符串转整形错误："+err.Error(), "E")
	}

	buffname := bufio.NewReader(os.Stdin)
	fmt.Print("请输入导出文件名称：")
	filename, _ := buffname.ReadString('\r')
	filename = strings.Replace(filename, "\r", "", -1)

	core.BatchCreateEmail(numint, filename)
}

/**
 * @description: 方法描述：批量空投nft
 * @Author: maxyang
 * @return {*}
 */
func airDropNft() {
	fmt.Println("空投NFT前，请确认主钱包账户里面有足够的主币能够支付链上gas费用")
	bufcontract := bufio.NewReader(os.Stdin)
	fmt.Print("请输入NFT合约地址：")
	contractename, _ := bufcontract.ReadString('\r')
	contractename = strings.Replace(contractename, "\r", "", -1)

	buffile := bufio.NewReader(os.Stdin)
	fmt.Print("请输入空投的excel文件名称：")
	filectename, _ := buffile.ReadString('\r')
	filectename = strings.Replace(filectename, "\r", "", -1)

	buffromadd := bufio.NewReader(os.Stdin)
	fmt.Print("请输入NFT转出钱包地址：")
	fromAdd, _ := buffromadd.ReadString('\r')
	fromAdd = strings.Replace(fromAdd, "\r", "", -1)

	bufprive := bufio.NewReader(os.Stdin)
	fmt.Print("请输入转出钱包私钥：")
	privekey, _ := bufprive.ReadString('\r')
	privekey = strings.Replace(privekey, "\r", "", -1)

	if contractename == "" || filectename == "" || fromAdd == "" || privekey == "" {
		fmt.Println("输入错误")
		os.Exit(0)
	}

	core.AirDropNft(contractename, filectename, fromAdd, privekey)
}
