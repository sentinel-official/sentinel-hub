package node

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/v12/x/node/keeper"
	v1keeper "github.com/sentinel-official/hub/v12/x/node/keeper/v1"
	v2keeper "github.com/sentinel-official/hub/v12/x/node/keeper/v2"
	v1types "github.com/sentinel-official/hub/v12/x/node/types/v1"
	v2types "github.com/sentinel-official/hub/v12/x/node/types/v2"
)

func RegisterServices(configurator sdkmodule.Configurator, cdc codec.BinaryCodec, k keeper.Keeper) {
	v1types.RegisterMsgServiceServer(configurator.MsgServer(), v1keeper.NewMsgServiceServer(k))
	v1types.RegisterQueryServiceServer(configurator.QueryServer(), v1keeper.NewQueryServiceServer(cdc, k))

	v2types.RegisterMsgServiceServer(configurator.MsgServer(), v2keeper.NewMsgServiceServer(k))
	v2types.RegisterQueryServiceServer(configurator.QueryServer(), v2keeper.NewQueryServiceServer(cdc, k))
}