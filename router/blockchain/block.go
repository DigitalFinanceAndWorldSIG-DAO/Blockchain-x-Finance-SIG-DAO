package blockchain

import (
	"fmt"
	"log"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
)

var instance *Blockchain
var contractAddr string = "0x8c7c77432a4b21d417b41ef5fea8212bf1f44b22"
var Cli *client.Client



//init会自动触发运行
func init() {
	//0. 读取解析配置文件
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	config := &configs[0]

	//1. 连接到节点
	cli, err := client.Dial(config)
	if err != nil {
		log.Fatal(err)
	}
	Cli = cli
	blockchain, err := NewBlockchain(common.HexToAddress(contractAddr), cli)
	if err != nil {
		log.Panic("failed to NewTask", err)
	}
	instance = blockchain
}


func Gethash() error {
	hash, error := instance.HashCode(Cli.GetCallOpts())
	if error != nil {
		return error
	} 
	fmt.Println(hash)
	return nil
}
func GetOwner() error {
	address, error := instance.Owner(Cli.GetCallOpts())
	if error != nil {
		return error
	} 
	fmt.Println(address)
	return nil
}


func GetCleartext(_cleartext string) error {
	_, _, err := instance.Getcleartext(Cli.GetTransactOpts(), _cleartext)
	if err != nil {
		return err
	}
	return nil
}

func GetCryphertext(_cryphertext string) error {
	_, _, err := instance.Getcleartext(Cli.GetTransactOpts(), _cryphertext)
	if err != nil {
		return err
	}
	return nil
}




