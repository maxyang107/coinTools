<!--
 * @Description: 
 * @Author: maxyang
 * @Date: 2022-01-07 17:11:23
 * @LastEditTime: 2022-01-11 09:43:45
 * @LastEditors: liutq
 * @Reference: 
-->
# coinTools
代码已开源
集成代币工具，主币归集，erc20代币归集，erc20代币空投，erc721nft空投。批量创建eth钱包。批量生成邮箱地址，如果对你也有帮助，请帮忙点一下星星，纯自己码的代码，欢迎大佬们指指点点，有其他工具的需求，欢迎大家提，我这边空了也会更新上来

# 使用方法
```
windows下可以直接下载coinTools，双击coinTools.exe
```
![image](https://user-images.githubusercontent.com/39045850/148506114-89b1352c-56c5-4964-8d85-771ecb97a2e2.png)

mac 及linux系统自己编译就可以了：

```
go mod download

go build -o coinTools ./

./coinTools

```

windows系统编译：

```
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o coinTools.exe ./

```

# 网络配置
当前默认是币安智能链主网

如果需要改网络的话，打开conf.json自行配置
```
{
    "ChanConn":  "https://bsc-dataseed1.binance.org",
	"GasLimmit": 210000
}
```

处理结果会记录在当前目录的log文件中

使用前务必看模板文件，对应文件才能正常解析，执行完成后，会在对应的excel文件中记录交易hash

![image](https://user-images.githubusercontent.com/39045850/148506463-d316a9f8-861f-4db5-8772-cf32d540b741.png)


注意事项
mac，linux环境下请将mian.go文件中的“ \r”替换成“\n”

