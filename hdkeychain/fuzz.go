package hdkeychain

import (
	"github.com/decred/dcrd/chaincfg"
	dcrd_hdkeychain "github.com/decred/dcrd/hdkeychain"
)

func Fuzz(input []byte) {
	key, err := dcrd_hdkeychain.NewKeyFromString(string(input))
	if err == nil {
		key.String()
	}
	if len(input) >= dcrd_hdkeychain.MinSeedBytes &&
		len(input) <= dcrd_hdkeychain.MaxSeedBytes {
		key, err = dcrd_hdkeychain.NewMaster(input, &chaincfg.MainNetParams)
		if err == nil {
			key.String()
		}
	}
}
