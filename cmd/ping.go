package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/shurcooL/githubv4"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func init() {
	rootCmd.AddCommand(pingCmd)
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "write message reminder for issues/prs forgotten since XX time",
	Long:  `ping authors with reminder`,
	Run: func(cmd *cobra.Command, args []string) {
		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		)
		httpClient := oauth2.NewClient(context.Background(), src)
		_ = githubv4.NewClient(httpClient)
		fmt.Printf("GITHUB_TOKEN:%s", viper.Get("GITHUB_TOKEN"))
	},
}
