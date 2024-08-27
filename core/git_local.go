package main

import (
    "github.com/go-git/go-git/v5"
    "github.com/go-git/go-git/v5/plumbing"
    "github.com/go-git/go-git/v5/plumbing/transport/http"
    "log"
    "os"
)

func PerformCloneAndInit(cloneUrl string) *git.Repository {
	repo, err := git.PlainClone("./", false, &git.CloneOptions{
		URL: cloneUrl,
		Auth: &http.BasicAuth{
			Username: os.Getenv("GIT_USERNAME"),// GIT_USERNAME
			Password: os.Getenv("GIT_ACCESS_TOKEN"), //GIT_ACCES_TOKEN
		},
		Progress: os.Stdout,
	})
	
	if err != nil {
		log.Fatal(err)
	}
	
	return repo
}

func CreateBranch(name string, repo *git.Repository) {
	branchRefName := plumbing.NewBranchReferenceName(name)
	w, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
	}
	err = w.Checkout(&git.CheckoutOptions{
		Branch: branchRefName,
		Force:  true,
		Create: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func GitPush(repo *git.Repository) {
	wt, _ := repo.Worktree()
	_, err := wt.Add("./")
	if err != nil {
		log.Fatal(err)
	}
	hash, err := wt.Commit("Initial Commit with template", &git.CommitOptions{AllowEmptyCommits: true})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s",hash.String())
	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &http.BasicAuth{
			Username: os.Getenv("GIT_USERNAME"),// GIT_USERNAME
			Password: os.Getenv("GIT_ACCESS_TOKEN"), //GIT_ACCES_TOKEN
		},
		Progress: os.Stdout,
	})
	if err != nil {
		log.Fatal(err)
	}
}