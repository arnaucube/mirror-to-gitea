# Mirror your github repositories to your gitea server

## Badges

[![image pulls](https://img.shields.io/docker/pulls/jaedle/mirror-to-gitea.svg)](https://cloud.docker.com/repository/docker/jaedle/mirror-to-gitea)
[![microbadger analysis](https://images.microbadger.com/badges/image/jaedle/mirror-to-gitea.svg)](https://microbadger.com/images/jaedle/mirror-to-gitea "Get your own image badge on microbadger.com")

## Description

`mirror-to-gitea` automatically creates repository mirrors from a Github User to your gitea server. The hard work (
continous mirroring) is done by Gitea on a continuous basis.

### Modes

There a different usage modes:

1. mirror **only public** repositories: No authentication on Github required. You may utilize an optional Github token
   to avoid rate limits
2. Mirror **public and private** repositories: Authentication for Github required.

### Prerequisites

#### Mandatory

- Something to mirror from Github (repositories from a Github user)
- An application token for a gitea instance

#### Optional

- A Github token (to *avoid rate limits* or to *mirror private repositories*)

### Examples

#### Mirror public repositories

A github user `github-user` has the public repositories `dotfiles` and `zsh-config`. Starting the script with a gitea
token for the account `gitea-user` will create the following mirror repositories:

- github.com/github-user/dotfiles &larr; some-gitea.url/gitea-user/dotfiles
- github.com/github-user/zsh-config &larr; some-gitea.url/zsh-config/dotfiles

The mirror settings are default by your gitea instance.

#### Mirror private repositories

You obtained a github token for `github-user` which has a public repository `public-example` a private
repository `private-example`. Running the mirroring process for `gitea-user` will create the following mirror
repositories:

- github.com/github-user/public-example &larr; some-gitea.url/gitea-user/public-example
- github.com/github-user/private-example &larr; some-gitea.url/zsh-config/private-example

The mirror settings are default by your gitea instance.

## Run public repository mirrors

### Parameters

#### Mirror public repositories

##### Mandatory

- `GITHUB_USERNAME` name of user with public repos should be mirrored
- `GITEA_URL` url of your gitea server
- `GITEA_TOKEN` token for your gitea user

##### Optional

- `MIRROR_REPOSITORIES` `PUBLIC` (optional, is default)
- `GITHUB_TOKEN` Github personal access token (optional to avoid rate limits)

```sh
docker container run \
 -d \
 --restart always \
 -e MIRROR_REPOSITORIES='PUBLIC' \
 -e GITHUB_USERNAME='github-user' \
 -e GITHUB_TOKEN='<optional-github-token>' \
 -e GITEA_URL=https://some-gitea.url \
 -e GITEA_TOKEN='<gitea-token>' \
 jaedle/mirror-to-gitea:latest
```

This will a spin up a docker container running infinite which will try to mirror all your repositories once every hour
to your gitea server.

## Run private repository mirrors

##### Mandatory

- `MIRROR_MODE` 'PRIVATE_AND_PUBLIC'
- `GITHUB_TOKEN` Github personal access token
- `GITEA_URL` url of your gitea server
- `GITEA_TOKEN` token for your gitea user

```sh
docker container run \
 -d \
 --restart always \
 -e MIRROR_MODE='PRIVATE_AND_PUBLIC' \
 -e GITHUB_TOKEN='<github-token>' \
 -e GITEA_URL=https://some-gitea.url \
 -e GITEA_TOKEN=please-exchange-with-token \
 jaedle/mirror-to-gitea:latest
```
