package models

import (
	"crypto/md5"
	"encoding/hex"
)

type Project struct {
	GithubRepo   string `json:"githubRepo"`
	GithubBranch string `json:"githubBranch"`
	GithubToken  string `json:"githubToken"`
}

func GetMD5Hash(proj Project) string {
	hash := md5.Sum([]byte(proj.GithubRepo + proj.GithubBranch))
	return hex.EncodeToString(hash[:])
}
