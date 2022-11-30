package config

import (
	"crypto/ecdsa"
)

const (
	Local = "http://127.0.0.1:8502"
	QA    = "http://172.22.42.159:8545"

	pk   = ""
	addr = "0x9fB29AAc15b9A4B7F17c3385939b007540f4d791"

	TokenHubContract = "0x0000000000000000000000000000000000001004"
	ChainId          = 714
)

type ExtAcc struct {
	RawKey string
	Key    *ecdsa.PrivateKey
}
