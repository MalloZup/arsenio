package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pingCmd)
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "write message reminder for issues/prs forgotten since XX time",
	Long:  `ping authors with reminder`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ping authors WIP")
	},
}
