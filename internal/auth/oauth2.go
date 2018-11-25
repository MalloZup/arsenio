package auth

import (
	"context"

	"github.com/shurcooL/githubv4"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// this function read via viper the token and initialize the gitHubClient
// with oauth2
func CreateGitHubClient() (*githubv4.Client, error) {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: viper.Get("GITHUB_TOKEN").(string)},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	gitHubClient := githubv4.NewClient(httpClient)

	return gitHubClient, nil
}
