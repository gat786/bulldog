package models

type ConfigFile struct {
	Config Config `yaml:"config"`
}

type Config struct {
	Resources        []GroupAndResourceNames `yaml:"resources"`
	Namespaces       []string                `yaml:"namespaces"`
	SaveFullManifest bool                    `yaml:"savefullmanifest"`
}
