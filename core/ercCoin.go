/*
 * @Description:ERC20代币归集
 * @Author: maxyang
 * @Date: 2022-01-06 17:01:58
 * @LastEditTime: 2022-01-10 13:41:20
 * @LastEditors: liutq
 * @Reference:
 */
package core

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/maxyang107/collectcoin/conn"
	"github.com/maxyang107/collectcoin/core/contract"
	"github.com/maxyang107/collectcoin/utils"
)

/**
 * @description: 方法描述：ERC20代币归集
 * @Author: maxyang
 * @return {*}
 * @param {string} contractAdd  合约地址
 * @param {string} filename  归集excel文件
 * @param {string} CollectionAddress  收款地址
 * @param {int} chanId 链id
 */
func CollectErc20Coin(contractAdd string, filename string, CollectionAddress string) {
	client := conn.Eclient
	xlsx, err := excelize.OpenFile(fmt.Sprintf("./%s.xlsx", filename))
	if err != nil {
		utils.WriteLog("读取excel文件错误："+err.Error(), "E")
		os.Exit(1)
	}
	rows := xlsx.GetRows("Sheet1")
	udi, err := contract.NewUdi(common.HexToAddress(contractAdd), client)
	if err != nil {
		utils.WriteLog("读取DUI错误："+err.Error(), "E")
		return
	}

	ToAddress := common.HexToAddress(CollectionAddress)
	var errnum int
	var succnum int
	var ziornum int
	for key, row := range rows {
		//查询余额
		blance, err := udi.BalanceOf(&bind.CallOpts{}, common.HexToAddress(row[0]))

		if err != nil {
			errnum++
			utils.WriteLog("获取ERC20代币错误，对应钱包地址："+row[0]+"已跳过该地址", "E")
			continue
		}
		//如果等于0，跳过
		if blance.Cmp(big.NewInt(0)) == 0 {
			ziornum++
			continue
		}

		//执行转账
		privateKey, err := crypto.HexToECDSA(row[1])
		if err != nil {
			utils.WriteLog("加密私钥错误："+err.Error(), "E")
			continue
		}

		nonce, err := client.NonceAt(context.Background(), common.HexToAddress(row[0]), nil)
		if err != nil {
			utils.WriteLog("获取nonce错误："+err.Error()+"对应钱包地址："+row[0], "E")
			continue
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			utils.WriteLog("获取gasprice错误："+err.Error()+"对应钱包地址："+row[0], "E")
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

		gasCost := utils.CalcGasCost(opts.GasLimit, opts.GasPrice)
		gasbalance, _ := client.BalanceAt(context.Background(), common.HexToAddress(row[0]), nil)

		if gasbalance.Cmp(gasCost) < 0 {
			utils.WriteLog("交易gas费用不足，对应钱包地址："+row[0]+"已跳过该地址", "T")
			xlsx.SetCellValue("Sheet1", fmt.Sprintf("C%d", key+1), "交易gas费用不足，对应钱包地址："+row[0]+"已跳过该地址")
			errnum++
			continue
		}

		tx, err := udi.Transfer(opts, ToAddress, blance)

		if err != nil {
			utils.WriteLog("获取ERC20代币错误，对应钱包地址："+row[0]+"已跳过该地址", "E")
			fmt.Println(err)
			xlsx.SetCellValue("Sheet1", fmt.Sprintf("C%d", key+1), "获取ERC20代币错误，对应钱包地址："+row[0]+"已跳过该地址")
			errnum++
			continue
		}
		succnum++
		utils.WriteLog("转出地址："+row[0]+"转入地址："+CollectionAddress+"交易hash："+tx.Hash().Hex(), "T")
		xlsx.SetCellValue("Sheet1", fmt.Sprintf("C%d", key+1), tx.Hash().Hex())
	}
	if err := xlsx.SaveAs(fmt.Sprintf("%s.xlsx", filename)); err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("归集任务执行完成，成功%d个，失败%d个，账户余额为0的账户%d个", succnum, errnum, ziornum))
	utils.WriteLog(fmt.Sprintf("归集任务执行完成，成功%d个，失败%d个，账户余额为0的账户%d个", succnum, errnum, ziornum), "T")
}
