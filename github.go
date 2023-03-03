package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/google/go-github/v50/github"
	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2"
)

const DEFAULT_REPOS_DIR = "gitscanner_repos_tmp"
const DEFAULT_OUTPUT_DIR = "findings"

/**
 * ScanGitHubOrganization scans a GitHub organization for vulnerabilities
 * @param cCtx cli context
 * @return error
 */
func ScanGitHubOrganization(cCtx *cli.Context) error {
	var token string = cCtx.String("token")
	var org string = cCtx.String("org")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	ctx := context.Background()
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repos, _, err := client.Repositories.List(ctx, org, nil)

	if err != nil {
		return err
	}

	for repo := range repos {
		println(repos[repo].GetCloneURL())
	}

	cloneAndScanAllRepositories(repos)

	return nil
}

func cloneAndScanAllRepositories(repos []*github.Repository) error {
	tmpDir := os.TempDir()
	tmpReposDir := filepath.Join(tmpDir, DEFAULT_REPOS_DIR)
	os.MkdirAll(tmpReposDir, os.ModePerm)
	defer os.RemoveAll(tmpReposDir)

	for repo := range repos {
		var repoPath string = filepath.Join(tmpReposDir, repos[repo].GetName())
		_, err := git.PlainClone(
			repoPath,
			false, // isBare
			&git.CloneOptions{
				URL:      repos[repo].GetCloneURL(),
				Progress: nil, // TODO: Make this configurable using verbose flag
			})
		if err != nil {
			return err
		}

		os.MkdirAll(DEFAULT_OUTPUT_DIR, os.ModePerm)

		out, err := exec.Command(
			"gitleaks",
			"detect",
			"--source", repoPath,
			"--exit-code", "0",
			"--report-path", filepath.Join(DEFAULT_OUTPUT_DIR, repos[repo].GetName()+"-report.json")).Output()
		if err != nil {
			println(err.Error())
			log.Fatal(err)
		}
		println(string(out))
	}

	return nil
}
