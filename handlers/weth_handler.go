package handlers

import (
	"context"
	"log"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

// Ethereum RPC endpoint
const rpcURL = "https://eth.llamarpc.com"

// WETH contract address on Ethereum Mainnet
const wethAddress = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

// Minimal ABI for the totalSupply function
const wethABI = `[{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`

type TotalSupplyResponse struct {
	TotalSupply string `json:"totalSupply"`
}

func GetWethTotalSupply(c *gin.Context) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
		// Use Gin's error handling
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Ethereum network"})
		return
	}
	defer client.Close()

	contractAddress := common.HexToAddress(wethAddress)
	parsedABI, err := abi.JSON(strings.NewReader(wethABI))
	if err != nil {
		log.Printf("Failed to parse contract ABI: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error (ABI parsing)"})
		return
	}

	data, err := parsedABI.Pack("totalSupply")
	if err != nil {
		log.Printf("Failed to pack data for totalSupply: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error (ABI packing)"})
		return
	}

	// Use ethereum.CallMsg type
	callMsg := ethereum.CallMsg{
		To:   &contractAddress,
		Data: data,
	}

	result, err := client.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		log.Printf("Failed to call contract: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to query contract"})
		return
	}

	var totalSupply *big.Int
	err = parsedABI.UnpackIntoInterface(&totalSupply, "totalSupply", result)
	if err != nil {
		log.Printf("Failed to unpack totalSupply result: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error (ABI unpacking)"})
		return
	}

	response := TotalSupplyResponse{
		TotalSupply: totalSupply.String(),
	}
	c.JSON(http.StatusOK, response)
}
