/*
 * @Description:批量归集主币excel版本
 * @Author: maxyang
 * @Date: 2022-01-06 10:52:51
 * @LastEditTime: 2022-01-07 17:07:16
 * @LastEditors: liutq
 * @Reference:
 */

package core

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/maxyang107/collectcoin/conn"
	"github.com/maxyang107/collectcoin/utils"
)

func CollectCoin(fileName string, CollectionAddress string) {
	client := conn.Eclient

	xlsx, err := excelize.OpenFile(fmt.Sprintf("./%s.xlsx", fileName))
	if err != nil {
		utils.Loger.Println("读取excel文件错误：" + err.Error())
		os.Exit(1)
	}
	rows := xlsx.GetRows("Sheet1")

	for _, row := range rows {
		account := common.HexToAddress(row[0])
		balance, err := client.BalanceAt(context.Background(), account, nil)
		if err != nil {
			utils.Loger.Println("读取账户余额错误：" + err.Error() + "对应钱包地址：" + row[0])
			continue
		}
		//如果等于0，跳过
		if balance.Cmp(big.NewInt(0)) == 0 {
			continue
		}

		privateKey, err := crypto.HexToECDSA(row[1])
		if err != nil {
			utils.Loger.Println("加密私钥错误：" + err.Error())
			continue
		}

		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			utils.Loger.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			continue
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			utils.Loger.Println("获取nonce错误：" + err.Error() + "对应钱包地址：" + row[0])
			continue
		}

		gasLimit := utils.ConfObj.GasLimmit
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			utils.Loger.Println("获取gasprice错误：" + err.Error() + "对应钱包地址：" + row[0])
			continue
		}

		gasCost := utils.CalcGasCost(gasLimit, gasPrice)
		value := balance.Sub(balance, gasCost)

		toAddress := common.HexToAddress(CollectionAddress)
		var data []byte
		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

		chainID, err := client.NetworkID(context.Background())

		if err != nil {
			utils.Loger.Println("获取链ID错误" + err.Error() + "对应钱包地址：" + row[0])
			continue
		}

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			utils.Loger.Println("交易签名错误" + err.Error() + "对应钱包地址：" + row[0])
			continue
		}

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			utils.Loger.Println("发送交易错误" + err.Error() + "对应钱包地址：" + row[0])
			continue
		}
		utils.Loger.Println("转出地址：" + row[0] + " 转入地址：" + CollectionAddress + " 交易hasd：" + signedTx.Hash().Hex())
	}
	fmt.Println("任务完成")
}
