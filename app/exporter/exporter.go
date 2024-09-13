package exporter

import (
	"fmt"
	"os"

	logrus "gat786/bulldog/log"
	yaml "github.com/ghodss/yaml"
)

// exists returns whether the given file or directory exists
func dirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func ExportYaml(namespaceData interface{}, fileName string) {
	marshalledData, err := yaml.Marshal(namespaceData)
	if err != nil {
		logrus.Error(err.Error())
	} else {
		marshalledString := string(marshalledData)
		output_dir := GetOutputDir()
		runOutputDir := fmt.Sprintf("%s/%s", output_dir, RUNTIME_STAMP)
		logrus.Info("Writing to directory: ", runOutputDir)
		checkExists, err := dirExists(runOutputDir)
		if err != nil {
			logrus.Error("Error checking directory: ", err.Error())
		}

		if !checkExists {
			err := os.MkdirAll(runOutputDir, 0755)
			if err != nil {
				logrus.Error("Error creating directory: ", err.Error())
			}
		}

		fileName := fmt.Sprintf("%s/%s.yaml", runOutputDir, fileName)
		file, err := os.Create(fileName)
		if err != nil {
			logrus.Error("Error creating file: ", err.Error())
		} else {
			defer file.Close()
			_, err := file.WriteString(marshalledString)
			if err != nil {
				logrus.Error("Error writing to file: ", err.Error())
			}
		}

	}
}
