package gh

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/costa86/github-automator/forms"

	"github.com/google/go-github/v63/github"
)

func GetClientAndContext() (*github.Client, context.Context) {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}

	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)

	return client, ctx
}

func GetRepos(ctx context.Context, client *github.Client) []string {
	var allRepos []*github.Repository
	var repoNames []string

	opt := &github.RepositoryListByAuthenticatedUserOptions{
		ListOptions: github.ListOptions{PerPage: 100},
	}
	for {
		repos, resp, err := client.Repositories.ListByAuthenticatedUser(ctx, opt)
		if err != nil {
			os.Exit(1)
		}
		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
	for _, v := range allRepos {
		repoNames = append(repoNames, *v.Name)
	}
	return repoNames
}

func CreateRepo(client *github.Client, ctx context.Context, repo forms.Repo) {

	r := &github.Repository{Name: &repo.Title, Private: &repo.Private, Description: &repo.Description, AutoInit: &repo.AutoInit}
	repoCreated, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repoCreated.GetName())
}

func DeleteRepo(ctx context.Context, client *github.Client, repo string) {

	user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		log.Fatal(err)
	}
	_, e := client.Repositories.Delete(ctx, *user.Login, repo)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	fmt.Printf("Successfully deleted  repo: %v\n", repo)
}
