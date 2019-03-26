## xtag

xtag command finds latest release tag which matched with given tag with `x` wild card (a.k.a. xtag).

### Installation

```shell
$ go get github.com/haya14busa/xtag/cmd/xtag
```

### Usage

```shell
$ xtag reviewdog/reviewdog 0.9.x
0.9.11
# Use GitHub API token to avoid API rate limit.
$ export GITHUB_API_TOKEN="xxx"
```
