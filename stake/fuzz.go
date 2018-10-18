package stake

import (
	dcrd_stake "github.com/decred/dcrd/blockchain/stake"
)

func Fuzz(input []byte) {
	outputs := []*dcrd_stake.MinimalOutput{
		{
			PkScript: input,
		},
		{
			PkScript: input,
		},
	}
	dcrd_stake.SStxStakeOutputInfo(outputs)
}
