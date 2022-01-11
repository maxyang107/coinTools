/*
 * @Description:
 * @Author: maxyang
 * @Date: 2022-01-06 18:04:09
 * @LastEditTime: 2022-01-10 14:53:13
 * @LastEditors: liutq
 * @Reference:
 */
package core

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/maxyang107/collectcoin/conn"
	"github.com/maxyang107/collectcoin/core/contract"
	"github.com/maxyang107/collectcoin/utils"
	"github.com/shopspring/decimal"
)

func AirDropErc20Coin(contractAdd string, filename string, fromAdd string, priveKey string) {
	client := conn.Eclient
	xlsx, err := excelize.OpenFile(fmt.Sprintf("./%s.xlsx", filename))
	if err != nil {
		fmt.Println("excel reder err:", err)
		utils.WriteLog("读取excel文件错误："+err.Error(), "E")
		os.Exit(1)
	}
	rows := xlsx.GetRows("Sheet1")
	udi, err := contract.NewUdi(common.HexToAddress(contractAdd), client)
	if err != nil {
		utils.WriteLog("读取DUI错误："+err.Error(), "E")
		return
	}

	privateKey, err := crypto.HexToECDSA(priveKey)
	if err != nil {
		utils.WriteLog("加密私钥错误："+err.Error(), "E")
		return
	}

	for key, row := range rows {

	LOOP:
		ToAddress := common.HexToAddress(row[0])

		nonce, err := client.NonceAt(context.Background(), common.HexToAddress(fromAdd), nil)
		if err != nil {
			utils.WriteLog("获取nonce错误："+err.Error()+"对应钱包地址："+row[0], "E")
			continue
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			utils.WriteLog("获取gasPrace错误："+err.Error()+"对应钱包地址："+row[0], "E")
			continue
		}

		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			utils.WriteLog("获取链ID错误"+err.Error()+"对应钱包地址："+row[0], "E")
			continue
		}
		opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			utils.WriteLog("组装options错误："+err.Error()+"对应钱包地址："+row[0], "E")
			continue
		}

		opts.GasLimit = utils.ConfObj.GasLimmit
		opts.Nonce = new(big.Int).SetInt64(int64(nonce))
		opts.GasPrice = gasPrice

		amount, err := decimal.NewFromString(row[1])
		if err != nil {
			utils.WriteLog("获取空投代币数量错误："+err.Error()+"对应钱包地址："+row[0], "E")
			continue
		}

		dec, err := udi.Decimals(nil)

		dectmp := decimal.NewFromFloat(math.Pow(10, float64(dec)))

		txAmount := amount.Mul(dectmp)
		airDropAmount, _ := new(big.Int).SetString(txAmount.String(), 10)

		fmt.Println(airDropAmount, ToAddress)

		tx, err := udi.Transfer(opts, ToAddress, airDropAmount)

		if err != nil {
			fmt.Println(err, "等待重试")
			time.Sleep(time.Second * time.Duration(rand.Int31n(5)))
			goto LOOP
		}

		utils.WriteLog("转出地址："+fromAdd+"转入地址："+row[0]+"交易hash："+tx.Hash().Hex(), "T")
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("C%d", key+1), tx.Hash().Hex())

	}
	if err := xlsx.SaveAs(fmt.Sprintf("%s.xlsx", filename)); err != nil {
		fmt.Println(err)
	}
	utils.WriteLog("ERC20空投任务完成", "T")
	fmt.Println("任务完成")
}
