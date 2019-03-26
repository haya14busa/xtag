package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v24/github"
	"github.com/haya14busa/xtag"
	"golang.org/x/oauth2"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: xtag [owner/repo] [xtag]")
	fmt.Fprintln(os.Stderr, "	xtag get latest GitHub released tag which matches given tag which containx 'x' as wild card (a.k.a xtag).")
	fmt.Fprintln(os.Stderr, "	Set GITHUB_API_TOKEN to avoid API limit.")
	fmt.Fprintln(os.Stderr, "Example:")
	fmt.Fprintln(os.Stderr, "	xtag reviewdog/reviewdog 0.9.x")
	fmt.Fprintln(os.Stderr, "	xtag golangci/golangci-lint 1.x")
	fmt.Fprintln(os.Stderr, "")
	fmt.Fprintln(os.Stderr, "GitHub: https://github.com/haya14busa/xtag")
}

var x = flag.Bool("test", false, "hoge")

func main() {
	flag.Usage = usage
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "xtag: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ownerAndRepo := flag.Arg(0)
	tag := flag.Arg(1)

	ss := strings.SplitN(ownerAndRepo, "/", 2)
	if len(ss) != 2 {
		return fmt.Errorf("invalid GitHub repository: %q", ownerAndRepo)
	}
	owner, repo := ss[0], ss[1]

	ctx := context.Background()

	tags, err := listTagsFromGitHubReleases(ctx, owner, repo)
	if err != nil {
		return err
	}

	latest, err := xtag.FindLatest(tag, tags)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stdout, latest)
	return nil
}

func listTagsFromGitHubReleases(ctx context.Context, owner, repo string) ([]string, error) {
	opt := &github.ListOptions{
		PerPage: 100,
	}
	cli := newGithubClient(ctx)
	tags, err := listTagsFromGitHubReleasesInternal(ctx, cli, owner, repo, opt, []string{})
	if err != nil {
		return nil, err
	}
	return tags, nil
}
func listTagsFromGitHubReleasesInternal(ctx context.Context, cli *github.Client, owner, repo string, opt *github.ListOptions, tags []string) ([]string, error) {
	releases, resp, err := cli.Repositories.ListReleases(ctx, owner, repo, opt)
	if err != nil {
		return nil, err
	}
	for _, r := range releases {
		if r.TagName != nil {
			tags = append(tags, *r.TagName)
		}
	}
	if resp.NextPage != 0 {
		opt.Page = resp.NextPage
		return listTagsFromGitHubReleasesInternal(ctx, cli, owner, repo, opt, tags)
	}
	return tags, nil
}

func newGithubClient(ctx context.Context) *github.Client {
	token := os.Getenv("GITHUB_API_TOKEN")
	var tc *http.Client
	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc = oauth2.NewClient(ctx, ts)
	}
	return github.NewClient(tc)
}
