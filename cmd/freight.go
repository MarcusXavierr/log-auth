package cmd

import (
	"fmt"

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

		if err != nil {
			fmt.Println("Error generating URL")
			panic(err)
		}

		fmt.Println("Generated url", url)
	},
}
