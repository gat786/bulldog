package models

type ExportedNamespaceData struct {
	NamespaceName string              `yaml:"namespaceName"`
	APIGroupMap   map[string]APIGroup `default:"{}",yaml:"apiGroup"`
}

type APIGroup struct {
	APIResourceMap map[string]APIResource `default:"{}",yaml:"apiResource"`
}

type APIResource struct {
	Resources []ResourceData `default:"[]",yaml:"resources"`
}

type ResourceData struct {
	ResourceName string `yaml:"resourceName"`
	Kind         string `yaml:"kind"`
	APIVersion   string `yaml:"apiVersion"`
	Manifest     string `default:""`
}
