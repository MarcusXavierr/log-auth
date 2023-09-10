package cmd

import (
	"os"

	"github.com/MarcusXavierr/log-auth/internal/IO"
	"github.com/MarcusXavierr/log-auth/internal/enum"
	"github.com/MarcusXavierr/log-auth/internal/login"
	"github.com/MarcusXavierr/log-auth/internal/stringgenerator"
	"github.com/spf13/cobra"
)

var freightIntelCmd = &cobra.Command{
	Use:   "freight",
	Short: "Gets the local url for freight intel",
	Run: func(cmd *cobra.Command, args []string) {
		url, err := stringgenerator.GeneratePBILocalUrl(login.MountLoginData(), enum.FreightIntel)
		IO.PrintGeneratedUrl(os.Stdout, url, err)
	},
}
