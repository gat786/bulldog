package kubernetes

import (
	"context"
	"fmt"

	"gat786/bulldog/exporter"
	"gat786/bulldog/models"

	logrus "gat786/bulldog/log"
	"github.com/ghodss/yaml"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func GetResources(loadedConfig models.Config) {
	dynamicClient := GetDynamicClient()

	apiGroupsAndResources := getApiGroupsAndResources(loadedConfig)
	availableNamespaces := getNamespacesToScrape(loadedConfig)

	scrapeConfig := make(map[string]interface{})
	scrapeConfig["namespaces"] = availableNamespaces
	scrapeConfig["resources"] = apiGroupsAndResources
	exporter.ExportYaml(scrapeConfig, "scrapeconfig")

	for _, namespaceName := range availableNamespaces {
		exportedNamespaceData := models.ExportedNamespaceData{
			NamespaceName: namespaceName,
			APIGroupMap:   make(map[string]models.APIGroup),
		}
		for _, apiGroupAndResources := range apiGroupsAndResources {
			var apiGroup models.APIGroup = models.APIGroup{
				APIResourceMap: make(map[string]models.APIResource),
			}
			apiGroupName := apiGroupAndResources.Group
			apiGroupVersion := apiGroupAndResources.Version
			apiGroupFullName := fmt.Sprintf("%s/%s", apiGroupName, apiGroupVersion)
			var resourcesToSearch []string
			if len(apiGroupAndResources.ResourceNames) == 0 {
				resourcesToSearch = getNamespacedResourcesForAPIGroup(apiGroupName)
			} else {
				resourcesToSearch = apiGroupAndResources.ResourceNames
			}
			for _, resourceName := range resourcesToSearch {
				var apiResource models.APIResource
				resourceGVRD := schema.GroupVersionResource{
					Group:    apiGroupAndResources.Group,
					Version:  apiGroupAndResources.Version,
					Resource: resourceName,
				}

				logrus.Infof("Searching Namespace: %s for Group %s", namespaceName, resourceGVRD)
				resourceManifests, err := dynamicClient.Resource(resourceGVRD).Namespace(namespaceName).List(context.TODO(), metav1.ListOptions{})
				var resourcesToBeExported []models.ResourceData
				if err != nil {
					logrus.Warn("Error fetching manifests for: ", resourceGVRD)
					logrus.Warn(err.Error())
				} else {
					if len(resourceManifests.Items) == 0 {
						logrus.Infof("No resources found for: %s in the namespace: %s", resourceGVRD, namespaceName)
					} else {
						for _, resourceManifest := range resourceManifests.Items {
							logrus.Info(resourceManifest.GetName())
							resourceData := models.ResourceData{
								ResourceName: resourceManifest.GetName(),
								Kind:         resourceManifest.GetKind(),
								APIVersion:   resourceManifest.GetAPIVersion(),
							}
							if loadedConfig.SaveFullManifest {
								marshalledBytes, err := yaml.Marshal(resourceManifest)
								if err != nil {
									logrus.Errorf("Error unmarshalling the manifest: %s", err.Error())
								} else {
									yamlManifest := string(marshalledBytes)
									resourceData.Manifest = yamlManifest
								}
							}
							resourcesToBeExported = append(resourcesToBeExported, resourceData)
						}
					}
				}

				if len(resourcesToBeExported) > 0 {
					apiResource = models.APIResource{
						Resources: resourcesToBeExported,
					}
					apiGroup.APIResourceMap[resourceName] = apiResource
				}
				exportedNamespaceData.APIGroupMap[apiGroupFullName] = apiGroup
			}
		}

		logrus.Info("Completed. Exporting the data")
		exporter.ExportYaml(exportedNamespaceData, namespaceName)
	}
}
