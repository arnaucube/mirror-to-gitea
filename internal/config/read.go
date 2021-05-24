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
	Gitea      GiteaConfig
	Github     GithubConfig
}

type GiteaConfig struct {
	GiteaUrl   string
	GiteaToken string
}

type GithubConfig struct {
	Username string
	Token    *string
}

func Read() (*Config, error) {
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

func readGiteaConfig() (GiteaConfig, error) {
	url, a := os.LookupEnv("GITEA_URL")
	if !a {
		return GiteaConfig{}, errors.New("missing mandatory parameter GITEA_URL, please specify your target gitea instance")
	}

	token, a := os.LookupEnv("GITEA_TOKEN")
	if !a {
		return GiteaConfig{}, errors.New("missing mandatory parameter GITEA_TOKEN, please specify your gitea application token")
	}

	return GiteaConfig{
		GiteaUrl:   url,
		GiteaToken: token,
	}, nil

}

func readGithubConfig() (GithubConfig, error) {
	username, present := os.LookupEnv("GITHUB_USERNAME")

	if !present {
		return GithubConfig{}, errors.New("")
	}

	var token *string = nil
	if val, hasToken := os.LookupEnv("GITHUB_TOKEN"); hasToken {
		token = &val
	}

	return GithubConfig{
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
