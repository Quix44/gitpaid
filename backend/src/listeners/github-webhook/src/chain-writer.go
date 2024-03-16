package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func payUser(_ context.Context, repoMetadata map[string]string, paymentAddress string, amount int, repo string) (*common.Hash, error) {
	pk := os.Getenv("HACKATHON_PK")
	if pk == "" {
		return nil, fmt.Errorf("private Key not set in environment")
	}

	// targetContractAddress := repoMetadata.PayeeAddress
	targetContractAddress := repoMetadata["contractAddress"]
	rpc := repoMetadata["rpc"]
	tokenAddress := repoMetadata["tokenAddress"]
	chainID := repoMetadata["chainID"]

	log.Printf("targetContractAddress: %s", targetContractAddress)
	log.Printf("rpc: %s", rpc)
	log.Printf("tokenAddress: %s", tokenAddress)
	log.Printf("chainID: %s", chainID)

	// Parse the contract ABI or we could load it from a file then parse it rather than env
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %v", err)
	}

	// Set up Ethereum client
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, fmt.Errorf("error creating Ethereum client: %v", err)
	}

	// Assuming amount is an int or int64, convert it to *big.Int
	decimalPlaces := big.NewInt(1_000_000_000_000_000_000)

	bigAmount := big.NewInt(int64(amount))

	inputData, err := parsedABI.Pack("pay", common.HexToAddress(tokenAddress), &repo, common.HexToAddress(paymentAddress), new(big.Int).Mul(bigAmount, decimalPlaces))
	if err != nil {
		return nil, fmt.Errorf("failed to pack input data: %v", err)
	}

	// Convert the contract address from a string to a common.Address
	contractAddress := common.HexToAddress(targetContractAddress)

	// Load the Private Key
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, fmt.Errorf("error loading private key: %v", err)
	}

	// Get the public key from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	// Get the address from the public key
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, fmt.Errorf("error getting nonce: %v", err)
	}

	// Get the gas price
	gasLimit := uint64(1000000) // We can lower this or do a suggestion call to get the gas limit
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting gas price: %v", err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, contractAddress, big.NewInt(int64(0)), gasLimit, gasPrice, inputData)

	chainId, err := strconv.Atoi(chainID)
	if err != nil {
		return nil, fmt.Errorf("error converting chain ID to int: %v", err)
	}
	// Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(int64(chainId))), privateKey)
	if err != nil {
		return nil, fmt.Errorf("error signing transaction: %v", err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, fmt.Errorf("error sending transaction: %v", err)
	}

	// Wait for the transaction to be confirmed
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return nil, fmt.Errorf("transaction wait error: %v", err)
	}

	// Check if the transaction failed
	if receipt.Status == types.ReceiptStatusFailed {
		return nil, fmt.Errorf("transaction failed: %v", receipt)
	}

	return &receipt.TxHash, nil
}

const contractABI = `[{"inputs":[{"internalType":"address","name":"_updaterAddress","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"uint256","name":"requested","type":"uint256"},{"internalType":"uint256","name":"available","type":"uint256"}],"name":"InsufficientFunds","type":"error"},{"inputs":[{"internalType":"uint256","name":"requestedTime","type":"uint256"},{"internalType":"uint256","name":"currentTime","type":"uint256"}],"name":"LockTimeNotMet","type":"error"},{"inputs":[{"internalType":"address","name":"owner","type":"address"}],"name":"OwnableInvalidOwner","type":"error"},{"inputs":[{"internalType":"address","name":"account","type":"address"}],"name":"OwnableUnauthorizedAccount","type":"error"},{"inputs":[],"name":"TransferFailed","type":"error"},{"inputs":[],"name":"Unauthorized","type":"error"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"token","type":"address"},{"indexed":false,"internalType":"string","name":"repository","type":"string"},{"indexed":false,"internalType":"address","name":"user","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"Funded","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"token","type":"address"},{"indexed":false,"internalType":"string","name":"repository","type":"string"},{"indexed":false,"internalType":"address","name":"to","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"PaymentMade","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"token","type":"address"},{"indexed":false,"internalType":"string","name":"repository","type":"string"},{"indexed":false,"internalType":"address","name":"user","type":"address"},{"indexed":false,"internalType":"uint256","name":"amount","type":"uint256"}],"name":"Withdrawn","type":"event"},{"inputs":[{"internalType":"address","name":"tokenAddress","type":"address"},{"internalType":"string","name":"repository","type":"string"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"fund","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"lockDuration","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"tokenAddress","type":"address"},{"internalType":"string","name":"repository","type":"string"},{"internalType":"address","name":"to","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"pay","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"","type":"address"},{"internalType":"string","name":"","type":"string"}],"name":"repositoryMap","outputs":[{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"uint256","name":"time","type":"uint256"},{"internalType":"address","name":"depositor","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"newUpdater","type":"address"}],"name":"setUpdaterAddress","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"totalFunded","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"totalPaid","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"updaterAddress","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"tokenAddress","type":"address"},{"internalType":"string","name":"repository","type":"string"},{"internalType":"uint256","name":"amount","type":"uint256"}],"name":"withdraw","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
