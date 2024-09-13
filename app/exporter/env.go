package exporter

import (
	"os"

	logrus "gat786/bulldog/log"
)

func GetOutputDir() string {
	outputs_dir, exists := os.LookupEnv(OUTPUTS_DIRECTORY_ENV_VAR)
	if !exists {
		logrus.Error("Outputs Directory not set, using default")
		outputs_dir = "./exported-data"
	}
	return outputs_dir
}
