/*
 * @Description:
 * @Author: maxyang
 * @Date: 2022-01-06 15:00:50
 * @LastEditTime: 2022-01-07 13:55:31
 * @LastEditors: liutq
 * @Reference:
 */
package core

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/maxyang107/collectcoin/utils"
)

func CreateWallet(num int, fileName string) {

	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	for i := 1; i < num+1; i++ {
		fmt.Println(fmt.Sprintf("正在生成第%d个钱包", i))
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}

		privateKeyBytes := crypto.FromECDSA(privateKey)

		privateKeyString := hexutil.Encode(privateKeyBytes)[2:]
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i), address)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i), privateKeyString)
		f.SetActiveSheet(index)
	}
	if err := f.SaveAs(fmt.Sprintf("%s.xlsx", fileName)); err != nil {
		fmt.Println(err)
	}
	utils.Loger.Println(fmt.Sprintf("钱包生成完成，共计生成了：%d个", num))
	fmt.Println("生成任务完成！")
}
