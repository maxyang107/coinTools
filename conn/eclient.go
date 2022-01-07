/*
 * @Description:
 * @Author: maxyang
 * @Date: 2022-01-06 10:53:23
 * @LastEditTime: 2022-01-07 14:11:09
 * @LastEditors: liutq
 * @Reference:
 */
package conn

import (
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/maxyang107/collectcoin/utils"
)

var Eclient *ethclient.Client

func init() {
	Eclient = newClient()
}

func newClient() *ethclient.Client {
	Eclient, err := ethclient.Dial(utils.ConfObj.ChanConn)
	if err != nil {
		utils.Loger.Println("主链连接错误错误：" + err.Error())
		os.Exit(0)
	}
	return Eclient
}
