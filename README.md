## xtag

xtag command finds latest release tag which matched with given tag with `x` wild card (a.k.a. xtag).

### Installation

```shell
# Install latest version
curl -sfL https://raw.githubusercontent.com/haya14busa/xtag/master/install.sh| sh -s

# binary will be in $(go env GOPATH)/bin/
curl -sfL https://raw.githubusercontent.com/haya14busa/xtag/master/install.sh| sh -s -- -b $(go env GOPATH)/bin [vX.Y.Z]

# or install it into ./bin/
curl -sfL https://raw.githubusercontent.com/haya14busa/xtag/master/install.sh| sh -s [vX.Y.Z]

# In alpine linux (as it does not come with curl by default)
wget -O - -q https://raw.githubusercontent.com/haya14busa/xtag/master/install.sh| sh -s [vX.Y.Z]
```

or you can use go get.

```shell
$ go get github.com/haya14busa/xtag/cmd/xtag
```

## Usage

```shell
$ xtag reviewdog/reviewdog 0.9.x
0.9.11
# Use GitHub API token to avoid API rate limit.
$ export GITHUB_API_TOKEN="xxx"
```

### Install v1.x golangci-lint CLI

```shell
$ curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin $(xtag golangci/golangci-lint v1.x)
golangci/golangci-lint info checking GitHub for tag 'v1.15.0'
golangci/golangci-lint info found version: 1.15.0 for v1.15.0/linux/amd64
golangci/golangci-lint info installed /home/haya14busa/bin/golangci-lint
```
