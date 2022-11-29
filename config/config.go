package config

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

const (
	Local = "http://127.0.0.1:8502"
	QA    = "http://172.22.42.159:8545"

	pk   = "9b28f36fbd67381120752d6172ecdcf10e06ab2d9a1367aac00cdcd6ac7855d3"
	addr = "0x9fB29AAc15b9A4B7F17c3385939b007540f4d791"

	TokenHubContract = "0x0000000000000000000000000000000000001004"
	ChainId          = 714
)

var TestAccount = ExtAcc{
	RawKey: pk,
	Addr:   common.HexToAddress(addr),
}

type ExtAcc struct {
	RawKey string
	Key    *ecdsa.PrivateKey
	Addr   common.Address
}
