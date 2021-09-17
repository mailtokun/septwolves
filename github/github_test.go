package github

import (
	"fmt"
	"os"
	"testing"
)

// TestClone calls github.Clone
func TestClone(t *testing.T) {
	os.Setenv("GITHUB_TOKEN", os.Getenv("GITHUB_TOKEN"))
	fmt.Println(os.Getenv("GITHUB_TOKEN"))
	if os.Getenv("GITHUB_TOKEN") == "" {
		t.Fatalf(`TestClone() = GITHUB_TOKEN is empty`)
	}
	os.Setenv("GITHUB_REPO", "https://github.com/mailtokun/rainbow")
	os.Setenv("GITHUB_BRANCH", "master")
	Clone()
	Pull()
}
