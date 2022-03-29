package localcmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func VerifyCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "verify",
		Short: "ver",
		Long:  "veri",
		RunE: func(cmd *cobra.Command, args []string) error {
			verifyNumbers()
			return nil
		},
	}
}

func verifyNumbers() {
	id := 100
	for {
		if id == 200 {
			break
		}
		fmt.Println(id, minQuantity(uint32(id)))
		id = id + 1

	}
}

func minQuantity(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
