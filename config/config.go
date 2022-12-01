package config

import (
	"crypto/ecdsa"
)

const (
	TokenHubContract = "0x0000000000000000000000000000000000001004"
	ChainId          = 714
)

type ExtAcc struct {
	RawKey string
	Key    *ecdsa.PrivateKey
}
