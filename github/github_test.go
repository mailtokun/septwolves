package github

import (
	"fmt"
	"github.com/mailtokun/yutu/models"
	"os"
	"testing"
)

// TestClone calls github.Clone
func TestClone(t *testing.T) {
	var proj models.Project
	proj.GithubRepo = "https://github.com/mailtokun/rainbow"
	proj.GithubToken = os.Getenv("GITHUB_TOKEN")

	fmt.Println(os.Getenv("GITHUB_TOKEN"))
	if os.Getenv("GITHUB_TOKEN") == "" {
		t.Fatalf(`TestClone() = GITHUB_TOKEN is empty`)
	}
	os.Setenv("GITHUB_REPO", "https://github.com/mailtokun/rainbow")
	os.Setenv("GITHUB_BRANCH", "master")
	//Clone(proj)
	fmt.Println(Pull(proj))
}
