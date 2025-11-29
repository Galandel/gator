package config

import (
	"os"
)

func getConfigFilePath() (string, error) {
	return os.UserHomeDir()
}
