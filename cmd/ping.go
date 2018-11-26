package cmd

import (
	"context"
	"fmt"

	gitHubAuth "github.com/MalloZup/arsenio/pkg/auth"
	"github.com/shurcooL/githubv4"
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
		gitHubClient, err := gitHubAuth.CreateGitHubClient()
		if err != nil {
			panic(err)
		}

		// TODO: do some stuff real later
		var query struct {
			Viewer struct {
				Login     githubv4.String
				CreatedAt githubv4.DateTime
			}
		}

		err = gitHubClient.Query(context.Background(), &query, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println("    Login:", query.Viewer.Login)
		fmt.Println("CreatedAt:", query.Viewer.CreatedAt)

	},
}
