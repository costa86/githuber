package forms

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

var accessible = os.Getenv("GITHUB_AUTOMATOR_ACCESSIBLE") != ""

type Action int

const (
	Delete Action = iota
	Create
)

func (s Action) String() string {
	switch s {
	case Delete:
		return "Delete repo "
	case Create:
		return "Create repo"
	default:
		return ""
	}
}

type Repo struct {
	Title       string
	Description string
	Private     bool
	AutoInit    bool
}

func minChar(s string) error {
	min := 5
	if len(s) < min {
		msg := fmt.Sprintf("Min %d chars in this field", min)
		return errors.New(msg)
	}
	return nil
}

func GetOperation() string {
	var operation string

	operationSelect := huh.NewSelect[string]().Options(huh.NewOptions(Create.String(), Delete.String())...).
		Title("Pick and action").Value(&operation)
	group1 := huh.NewGroup(operationSelect)

	form := huh.NewForm(group1).WithAccessible(accessible)
	err := form.Run()

	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
	return operation
}

func DeleteRepoForm(repos []string) string {
	var selectedRepo string
	repoSelector := huh.NewSelect[string]().Options(huh.NewOptions(repos...)...).
		Title(fmt.Sprintf("Pick a repo (%d)", len(repos))).Value(&selectedRepo)
	group1 := huh.NewGroup(repoSelector)
	form := huh.NewForm(group1).WithAccessible(accessible)
	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}
	return selectedRepo
}

func DeleteRepoConfirmForm(repo string) bool {
	var delete bool
	confirm := huh.NewConfirm().Title(fmt.Sprintf("Are you sure you want to delete %s?", repo)).Value(&delete)
	group1 := huh.NewGroup(confirm)

	form := huh.NewForm(group1).WithAccessible(accessible)
	err := form.Run()

	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	return delete
}

func CreateRepoForm() Repo {
	var repo Repo

	title := huh.NewInput().Title("Title").Value(&repo.Title).Validate(minChar)
	description := huh.NewInput().Title("Description").Value(&repo.Description)
	private := huh.NewConfirm().Title("Private").Value(&repo.Private)
	autoInit := huh.NewConfirm().Title("Auto-init").Value(&repo.AutoInit).Description("Create an empty README.md")

	group1 := huh.NewGroup(title, description, private, autoInit)

	form := huh.NewForm(group1).WithAccessible(accessible)
	err := form.Run()

	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	return repo
}
