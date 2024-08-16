package main

import (
	"context"

	"github.com/costa86/github-automator/forms"
	"github.com/costa86/github-automator/gh"

	"github.com/google/go-github/v63/github"
)

func createRepoWrapper(ctx context.Context, client *github.Client) {
	repo := forms.CreateRepoForm()
	repoUrl := gh.CreateRepo(client, ctx, repo)
	forms.RunGitCommands(repo.Folder, repoUrl)
}

func deleteRepoWrapper(ctx context.Context, client *github.Client) {
	repos := gh.GetRepos(ctx, client)
	selectedRepo := forms.DeleteRepoForm(repos)
	confirmation := forms.DeleteRepoConfirmForm(selectedRepo)

	if confirmation {
		gh.DeleteRepo(ctx, client, selectedRepo)
	}
}

func main() {
	client, context := gh.GetClientAndContext()

	switch forms.GetOperation() {

	case forms.Delete.String():
		deleteRepoWrapper(context, client)

	case forms.Create.String():
		createRepoWrapper(context, client)
	}

}
