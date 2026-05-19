package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	createWallet()
}

func createWallet() {
	//creates a new wallet
	wallet := solana.NewWallet()
	fmt.Printf("My wallet address: %s\n", wallet.PublicKey().String())

	//connect to the DevNet network
	client := rpc.New(rpc.DevNet_RPC)
	defer func() {
		if err := client.Close(); err != nil {
			log.Println("error closing client:", err)
		}
	}()
	result, err := client.GetBalance(context.Background(), solana.MustPublicKeyFromBase58("ASr7BW5o7PdQfrg6zjtzT8Y3i4GwNukfoSnu196Y6u2U"), rpc.CommitmentFinalized)
	if err != nil {
		log.Fatal(err)
	}
	balanceInSol := float64(result.Value) / 1_000_000_000
	fmt.Printf("Your Balance: %f SOL\n", balanceInSol)
}
