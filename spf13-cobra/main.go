package main

import (
	"fmt"

	localcmd "spf13-cobra/local_cmd"

	"github.com/spf13/cobra"
)

func main() {
	fmt.Println("server start")
	mainCmd := &cobra.Command{Use: "spf13-cobra"}
	mainCmd.AddCommand(localcmd.VersionCMD())
	mainCmd.AddCommand(localcmd.VerifyCMD())
	err := mainCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
