package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"inbox/contracts"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}

	// 合约地址
	address := common.HexToAddress("0x7089CD028220F44C6D5a92AF0cFdb4480Aab6Ed6")
	instance, err := contracts.NewContracts(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")

	// 调用合约账户私钥
	privateKey, err := crypto.HexToECDSA("a6da2e26b3e5c10b5c0c89974d9f9dbe99b7cee438ae588f9b5b6a989e0b930f")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	result, err := instance.Message(nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Before:", string(result[:]))

	tx, err := instance.SetMessage(auth, "hello"+strconv.Itoa(rand.Intn(10000)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	result, err = instance.Message(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("After:", string(result[:]))
}
