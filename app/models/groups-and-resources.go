package models

import ()

type GroupAndResourceNames struct {
	Group         string   `yaml:"group"`
	Version       string   `yaml:"version"`
	ResourceNames []string `yaml:"resourcenames"`
}
