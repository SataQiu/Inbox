#
使用 go-ethereum 操作简单智能合约

## 前置准备
安装 Ganache 运行本地私有链
安装 solc、apigen 工具

## 编译合约

```bash
cd contracts
solcjs --abi Inbox.sol
solcjs --bin Inbox.sol
abigen --bin=Inbox_sol_Inbox.bin --abi=Inbox_sol_Inbox.abi --pkg=contracts --out=inbox.go
```

## 部署合约

替换程序内的地址，为本地私有链地址，默认为 http://127.0.0.1:7545
替换程序内的账户私钥，为 Ganache 其中一个账户的私钥

```bash
go run cmd/deploy.go
```

## 加载并调用合约

替换程序内的地址，为本地私有链地址，默认为 http://127.0.0.1:7545
替换程序内合约地址，为执行合约后产生的合约地址
替换程序内的调用合约账户私钥，为 Ganache 其中一个账户的私钥

调用合约
```bash
go run cmd/call.go
```
