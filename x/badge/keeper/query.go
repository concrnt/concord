package keeper

import (
	"github.com/concrnt/concord/x/badge/types"
)

var _ types.QueryServer = Keeper{}
