package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mkXultra/starport_first_scaffold/testutil/keeper"
	"github.com/mkXultra/starport_first_scaffold/x/starportfirstscaffold/keeper"
	"github.com/mkXultra/starport_first_scaffold/x/starportfirstscaffold/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.StarportfirstscaffoldKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
