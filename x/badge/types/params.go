package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	DefaultCreateSeriesCost = 1000000 // 1 token
	DefaultMintBadgeCost    = 100000  // 0.1 token
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(createSeriesCost, mintBadgeCost uint64) Params {
	return Params{
		CreateSeriesCost: createSeriesCost,
		MintBadgeCost:    mintBadgeCost,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		CreateSeriesCost: DefaultCreateSeriesCost,
		MintBadgeCost:    DefaultMintBadgeCost,
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}
