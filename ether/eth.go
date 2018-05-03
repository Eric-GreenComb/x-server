package ether

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

// EthClient EthClient
var EthClient *ethclient.Client

// LoadEthClient LoadEthClient
func LoadEthClient() {
	EthClient, err := ethclient.Dial("http://localhost:8540")
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
	}
	fmt.Println(EthClient)
}

// GetEthClient GetEthClient
func GetEthClient() *ethclient.Client {
	client, err := ethclient.Dial("http://localhost:8540")
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client:", err)
	}
	return client
}
