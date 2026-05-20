package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

type WalletFile struct {
	SecretKey []byte `json:"secretKey"`
}

func main() {
	createWallet()
}

func loadOrCreateWallet() (*solana.Wallet, error) {
	file, err := os.Open("wallet.json")
	if err != nil {

		wallet := solana.NewWallet()
		createFile, err := os.Create("wallet.json")
		if err != nil {
			return nil, err
		}
		defer func() {
			if err := createFile.Close(); err != nil {
				log.Println("error closing client:", err)
			}
		}()

		var walletFile WalletFile
		privateKey := wallet.PrivateKey

		walletFile.SecretKey = privateKey
		encoder := json.NewEncoder(createFile)
		err = encoder.Encode(walletFile)
		if err != nil {
			return nil, err
		}
		return wallet, nil
	} else {
		var walletFile WalletFile
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&walletFile)
		if err != nil {
			return nil, err
		}
		wallet := &solana.Wallet{
			PrivateKey: solana.PrivateKey(walletFile.SecretKey),
		}
		return wallet, nil
	}
}

func createWallet() {
	//creates a new wallet
	wallet, err := loadOrCreateWallet()
	if err != nil {
		fmt.Println("Error: creating a new wallet", err)
	}
	fmt.Printf("My wallet address: %s\n", wallet.PublicKey().String())

	//connect to the DevNet network
	client := rpc.New(rpc.DevNet_RPC)
	defer func() {
		if err := client.Close(); err != nil {
			log.Println("error closing client:", err)
		}
	}()
	//getbalance
	result, err := client.GetBalance(context.Background(), wallet.PublicKey(),
		rpc.CommitmentFinalized)
	if err != nil {
		log.Fatal(err)
	}
	//convert Lamport to SOL
	balanceInSol := float64(result.Value) / 1_000_000_000
	fmt.Printf("Your Balance: %f SOL\n", balanceInSol)
}
