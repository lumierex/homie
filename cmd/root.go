package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	RootCmd.AddCommand(&EnvCmd)
	RootCmd.AddCommand(&QRCodeCmd)
}
