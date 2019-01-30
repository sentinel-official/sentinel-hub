package main

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmTypes "github.com/tendermint/tendermint/types"

	app "github.com/ironman0x7b2/sentinel-sdk/apps/vpn"
	vpnCli "github.com/ironman0x7b2/sentinel-sdk/apps/vpn/cli"
)

func main() {
	cdc := app.MakeCodec()

	config := csdkTypes.GetConfig()
	config.SetBech32PrefixForAccount(csdkTypes.Bech32PrefixAccAddr, csdkTypes.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(csdkTypes.Bech32PrefixValAddr, csdkTypes.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(csdkTypes.Bech32PrefixConsAddr, csdkTypes.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "vpnd",
		Short:             "VPN Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(vpnCli.InitCmd(ctx, cdc))
	rootCmd.AddCommand(vpnCli.CollectGenTxsCmd(ctx, cdc))
	rootCmd.AddCommand(vpnCli.TestnetFilesCmd(ctx, cdc))
	rootCmd.AddCommand(vpnCli.GenTxCmd(ctx, cdc))
	rootCmd.AddCommand(vpnCli.AddGenesisAccountCmd(ctx, cdc))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	executor := cli.PrepareBaseCmd(rootCmd, "SH", app.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db tmDB.DB, traceStore io.Writer) abciTypes.Application {
	return app.NewVPN(
		logger, db, traceStore, true,
		baseapp.SetPruning(store.NewPruningOptions(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
	)
}

func exportAppStateAndTMValidators(logger log.Logger, db tmDB.DB, traceStore io.Writer, height int64,
	forZeroHeight bool) (json.RawMessage, []tmTypes.GenesisValidator, error) {
	if height != -1 {
		vpn := app.NewVPN(logger, db, traceStore, false)
		err := vpn.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return vpn.ExportAppStateAndValidators(forZeroHeight)
	}
	vpn := app.NewVPN(logger, db, traceStore, true)
	return vpn.ExportAppStateAndValidators(forZeroHeight)
}