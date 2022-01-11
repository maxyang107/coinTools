/*
 * @Description:批量空投nft
 * @Author: maxyang
 * @Date: 2022-01-10 14:01:53
 * @LastEditTime: 2022-01-10 17:15:15
 * @LastEditors: liutq
 * @Reference:
 */
package core

import (
	"context"
	"fmt"
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

func AirDropNft(contractAdd string, filename string, fromAdd string, priveKey string) {
	client := conn.Eclient
	xlsx, err := excelize.OpenFile(fmt.Sprintf("./%s.xlsx", filename))
	if err != nil {
		fmt.Println("excel reder err:", err)
		utils.WriteLog("读取excel文件错误："+err.Error(), "E")
		os.Exit(1)
	}
	rows := xlsx.GetRows("Sheet1")
	erc721, err := contract.NewNft(common.HexToAddress(contractAdd), client)
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

		tokenIdstr, err := decimal.NewFromString(row[1])
		if err != nil {
			utils.WriteLog("获取空投Nft tokenId错误："+err.Error()+"对应钱包地址："+row[0], "E")
			continue
		}

		tokenId, _ := new(big.Int).SetString(tokenIdstr.String(), 10)

		tx, err := erc721.SafeTransferFrom(opts, common.HexToAddress(fromAdd), ToAddress, tokenId)

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
	utils.WriteLog("NFT空投任务完成", "T")
	fmt.Println("任务完成")
}
