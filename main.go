package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v63/github"
)

var (
	name        = flag.String("name", "", "Name of repo to create in authenticated user's GitHub account.")
	description = flag.String("description", "", "Description of created repo.")
	private     = flag.Bool("private", false, "Will created repo be private.")
	autoInit    = flag.Bool("auto-init", false, "Pass true to create an initial commit with empty README.")
	delete      = flag.Bool("delete", false, "Will delete a repo.")
)

func main() {
	flag.Parse()

	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}

	if *name == "" {
		log.Fatal("No name: New repos must be given a name")
	}

	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(token)

	if *delete {
		_, e := client.Repositories.Delete(ctx, "costa86", *name)
		if e != nil {
			fmt.Println(e)
			return
		}
		fmt.Printf("Successfully deleted  repo: %v\n", *name)
		return
	}

	r := &github.Repository{Name: name, Private: private, Description: description, AutoInit: autoInit}
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo: %v\n", repo.GetName())
}
