package localcmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func VersionCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "v",
		Long:  "lv",
		RunE: func(cmd *cobra.Command, args []string) error {
			printversion()
			return nil
		},
	}
}

func printversion() {
	fmt.Println("version 1")
}
