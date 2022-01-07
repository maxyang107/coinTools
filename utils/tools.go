/*
 * @Description:工具
 * @Author: maxyang
 * @Date: 2022-01-07 12:04:02
 * @LastEditTime: 2022-01-07 12:04:03
 * @LastEditors: liutq
 * @Reference:
 */
package utils

import "math/big"

//计算交易gas话费
func CalcGasCost(gasLimit uint64, gasPrice *big.Int) *big.Int {
	gasLimitBig := big.NewInt(int64(gasLimit))
	return gasLimitBig.Mul(gasLimitBig, gasPrice)
}
