package main

import (
	"context"
	"flag"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"test-crosschain-transfer/abi"
	"test-crosschain-transfer/config"
)

var (
	rpcUrl     = flag.String("rpc-url", "", "rpc url")
	amount     = flag.Int64("amount", 100, "transfer amount(BNB)")
	recipient  = flag.String("to", "", "recipient address")
	privateKey = flag.String("private-key", "9b28f36fbd67381120752d6172ecdcf10e06ab2d9a1367aac00cdcd6ac7855d3", "sender's private key")
)

func main() {
	flag.Parse()
	client, err := ethclient.Dial(*rpcUrl)
	if err != nil {
		log.Fatal(err)
	}

	account := config.ExtAcc{RawKey: *privateKey}
	account.Key, _ = crypto.HexToECDSA(account.RawKey)
	tokenHub, err := abi.NewTokenHub(common.HexToAddress(config.TokenHubContract), client)
	if err != nil {
		log.Fatal(err)
	}

	ops, _ := bind.NewKeyedTransactorWithChainID(account.Key, big.NewInt(config.ChainId))
	bnbAddr := common.HexToAddress("0x0000000000000000000000000000000000000000")
	toAddr := common.HexToAddress(*recipient)
	amt := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(*amount))
	expiredTime := time.Now().Unix() + 300
	ops.Value = new(big.Int).Add(new(big.Int).Mul(big.NewInt(1e17), big.NewInt(1)), amt)

	tx, err := tokenHub.TransferOut(ops, bnbAddr, toAddr, amt, uint64(expiredTime))
	if err != nil {
		log.Fatal("Error transfer to BC:", err)
	}
	var rc *types.Receipt
	for i := 0; i < 180; i++ {
		rc, err = client.TransactionReceipt(context.Background(), tx.Hash())
		if err == nil && rc.Status != 0 {
			break
		} else if rc != nil && rc.Status == 0 {
			log.Fatal("Transfer to BC failed")
		}
		time.Sleep(100 * time.Millisecond)
	}
	if rc == nil {
		log.Fatal("Transfer to BC failed")
	} else {
		log.Printf("Transfer to %s %d BNB succeed", *recipient, *amount)
	}
}
