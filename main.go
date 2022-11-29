package main

import (
	"context"
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

func main() {
	client, err := ethclient.Dial(config.QA)
	if err != nil {
		log.Fatal(err)
	}

	account := config.TestAccount
	account.Key, _ = crypto.HexToECDSA(account.RawKey)
	tokenHub, _ := abi.NewTokenHub(common.HexToAddress(config.TokenHubContract), client)

	ops, _ := bind.NewKeyedTransactorWithChainID(account.Key, big.NewInt(config.ChainId))
	ops.Value = new(big.Int).Mul(big.NewInt(1e17), big.NewInt(1001))
	bnbAddr := common.HexToAddress("0x0000000000000000000000000000000000000000")
	recipient := common.HexToAddress("0x0cdce3d8d17c0553270064cee95c73f17534d5a0")
	amount := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100))
	expiredTime := time.Now().Unix() + 300
	tx, err := tokenHub.TransferOut(ops, bnbAddr, recipient, amount, uint64(expiredTime))
	if err != nil {
		log.Fatal("Error transfer to BC:", err)
	}
	var rc *types.Receipt
	for i := 0; i < 180; i++ {
		rc, err = client.TransactionReceipt(context.Background(), tx.Hash())
		if err == nil && rc.Status != 0 {
			break
		} else if rc != nil && rc.Status == 0 {
			log.Fatal("Register relayer failed")
		}
		time.Sleep(100 * time.Millisecond)
	}
	if rc == nil {
		log.Fatal("Transfer to BC failed")
	} else {
		log.Printf("Transfer to %s 100 BNB succeed", "0x0cdce3d8d17c0553270064cee95c73f17534d5a0")
	}
}
