package cmd

import (
	"fmt"
	"os"

	"github.com/MarcusXavierr/log-auth/internal/utils"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "log-auth",
	Long: `Log auth é uma ferramenta CLI que visa facilitar o login na plataforma logcomex para fins de desenvolvimento local`,
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
