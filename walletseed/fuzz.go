package hdkeychain

import (
	dcrwallet_walletseed "github.com/decred/dcrwallet/walletseed"
)

func Fuzz(input []byte) {
	dcrwallet_walletseed.EncodeMnemonicSlice(input)
	dcrwallet_walletseed.EncodeMnemonic(input)
	dcrwallet_walletseed.DecodeUserInput(string(input))
}
