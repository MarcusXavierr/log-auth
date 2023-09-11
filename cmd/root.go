package cmd

import (
	"fmt"
	"os"

	"github.com/MarcusXavierr/log-auth/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "log-auth",
	Long: `Log auth is a CLI tool aimed at simplifying the login process to the Logcomex platform for local development purposes`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return utils.InitializeViper()
	},
}

func init() {
	rootCmd.AddCommand(shipmentCmd, freightIntelCmd)
	rootCmd.PersistentFlags().StringP("desired-platform", "d", "others", "")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
