package config

import (
	"errors"
	"fmt"
	"os"
)

const MirrorModePublic = "PUBLIC"
const MirrorModePrivateAndPublic = "PRIVATE_AND_PUBLIC"

type Config struct {
	MirrorMode string
	Gitea      Gitea
	Github     Github
}

type Gitea struct {
	Url   string
	Token string
}

type Github struct {
	Username string
	Token    *string
}

func readGiteaConfig() (Gitea, error) {
	url, a := os.LookupEnv("GITEA_URL")
	if !a {
		return Gitea{}, errors.New("missing mandatory parameter GITEA_URL, please specify your target gitea instance")
	}

	token, a := os.LookupEnv("GITEA_TOKEN")
	if !a {
		return Gitea{}, errors.New("missing mandatory parameter GITEA_TOKEN, please specify your gitea application token")
	}

	return Gitea{
		Url:   url,
		Token: token,
	}, nil

}

func readGithubConfig() (Github, error) {
	username, present := os.LookupEnv("GITHUB_USERNAME")

	if !present {
		return Github{}, errors.New("")
	}

	var token *string = nil
	if val, hasToken := os.LookupEnv("GITHUB_TOKEN"); hasToken {
		token = &val
	}

	return Github{
		Username: username,
		Token:    token,
	}, nil
}

func readMirrorMode() (string, error) {
	input, present := os.LookupEnv("MIRROR_MODE")

	if !present {
		return MirrorModePublic, nil
	}

	switch input {
	case MirrorModePublic, MirrorModePrivateAndPublic:
		return input, nil
	default:
		return "", fmt.Errorf("unknown mirror mode %s, please specify a valid mirror mode: PUBLIC, PRIVATE_AND_PUBLIC", input)
	}

}

type Reader struct {
}

func (r Reader) Read() (*Config, error) {
	mirrorMode, err := readMirrorMode()
	if err != nil {
		return nil, err
	}

	gitea, err := readGiteaConfig()
	if err != nil {
		return nil, err
	}

	github, err := readGithubConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		MirrorMode: mirrorMode,
		Gitea:      gitea,
		Github:     github,
	}, nil
}

func NewReader() *Reader {
	return &Reader{}
}
