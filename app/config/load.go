package config

import (
	"os"

	logrus "gat786/bulldog/log"
	"gat786/bulldog/models"
	yaml "github.com/ghodss/yaml"
)

func LoadConfig() models.Config {
	defaultConfig := models.Config{
		Resources:        make([]models.GroupAndResourceNames, 0),
		Namespaces:       make([]string, 0),
		SaveFullManifest: false,
	}
	configFilePath, exists := os.LookupEnv(CONFIG_FILE_PATH_ENV_VAR)
	if !exists {
		logrus.Info("CONFIG_FILE_PATH not provided, returning default config", configFilePath)
		return defaultConfig
	} else {
		configFileBytes, err := os.ReadFile(configFilePath)
		if err != nil {
			logrus.Errorf("Error Reading config file %s", err.Error())
			logrus.Errorf("Using Default config file instead")
			return defaultConfig
		}
		var loadedConfigFile models.ConfigFile
		err = yaml.Unmarshal(configFileBytes, &loadedConfigFile)
		if err != nil {
			logrus.Error("Error unmarshalling Config File make sure they are in correct format")
			logrus.Errorf("%s", err.Error())
			logrus.Error("Using Default Config File")
			return defaultConfig
		}
		logrus.Debugf("Successfully Loaded supplied configuration at path: %s", configFilePath)
		logrus.Debugf("Loaded Configuration: %+v", loadedConfigFile)
		return loadedConfigFile.Config
	}
}
