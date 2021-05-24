package app

import "github.com/jaedle/mirror-to-gitea/internal/config"

type ConfigReader interface {
	Read() (*config.Config, error)
}
