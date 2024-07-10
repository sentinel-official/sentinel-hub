package cli

import (
	"github.com/spf13/cobra"
)

func GetQueryCommands() []*cobra.Command {
	return []*cobra.Command{
		queryLease(),
		queryLeases(),
		queryParams(),
	}
}

func GetTxCommands() []*cobra.Command {
	cmd := &cobra.Command{
		Use:   "lease",
		Short: "Lease module sub-commands",
	}

	cmd.AddCommand(
		txStart(),
		txUpdate(),
		txRenew(),
		txEnd(),
	)

	return []*cobra.Command{cmd}
}
