package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

func main() {
	client := rpc.New(rpc.DevNet_RPC)

	defer func() {
		if err := client.Close(); err != nil {
			log.Println("error closing client: ", err)
		}
	}()

	sig, err := client.GetSignaturesForAddress(context.Background(),
		solana.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb"))
	if err != nil {
		log.Fatal(err)
	}

	for _, sign := range sig {
		status := "success"
		if sign.Err != nil {
			status = "failed"
		}
		fmt.Printf("transaction ID: %s\nslot number: %d\ntimestamp: %s\nstatus : %v\n\n",
			sign.Signature, sign.Slot, sign.BlockTime, status)
	}
}
