package config

import (
	"errors"
	"os"
)

const MirrorModePublic = "PUBLIC"
const MirrorModePrivateAndPublic = "PRIVATE_AND_PUBLIC"

type Config struct {
	MirrorMode string
	Gitea      GiteaConfig
}

type GiteaConfig struct {
	GiteaUrl   string
	GiteaToken string
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

	return &Config{
		MirrorMode: mirrorMode,
		Gitea:      gitea,
	}, nil
}

func readGiteaConfig() (GiteaConfig, error) {
	url, a := os.LookupEnv("GITEA_URL")
	if !a {
		return GiteaConfig{}, errors.New("")
	}

	token, a := os.LookupEnv("GITEA_TOKEN")
	if !a {
		return GiteaConfig{}, errors.New("")
	}

	return GiteaConfig{
		GiteaUrl:   url,
		GiteaToken: token,
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
		return "", errors.New("")
	}

}
