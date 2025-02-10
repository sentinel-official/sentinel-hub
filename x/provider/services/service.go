package services

import (
	sdkmodule "github.com/cosmos/cosmos-sdk/types/module"

	"github.com/sentinel-official/hub/v12/x/provider/keeper"
	"github.com/sentinel-official/hub/v12/x/provider/services/v2"
	"github.com/sentinel-official/hub/v12/x/provider/services/v3"
	v2types "github.com/sentinel-official/hub/v12/x/provider/types/v2"
	v3types "github.com/sentinel-official/hub/v12/x/provider/types/v3"
)

func RegisterServices(configurator sdkmodule.Configurator, k keeper.Keeper) {
	v3types.RegisterMsgServiceServer(configurator.MsgServer(), v3.NewMsgServiceServer(k))

	v2types.RegisterQueryServiceServer(configurator.QueryServer(), v2.NewQueryServiceServer(k))
}
