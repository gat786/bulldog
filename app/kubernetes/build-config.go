package kubernetes

import (
	"context"
	// "fmt"
	"os"
	"strings"

	"gat786/bulldog/models"

	logrus "gat786/bulldog/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getNamespacesToScrape(loadedConfig models.Config) []string {
	clientset := GetClient()
	availableNamespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logrus.Errorf("Error fetching namespaces: %s", err.Error())
		os.Exit(1)
	}
	var namespaces []string

	if len(loadedConfig.Namespaces) > 0 {
		logrus.Infof("Found Specified namespaces in config file provided. Only these namespaces would be scraped: %+v", loadedConfig.Namespaces)
		return loadedConfig.Namespaces
	}
	logrus.Info("No namespaces found in provided config, will scrape all the namespaces")
	for _, namespace := range availableNamespaces.Items {
		namespaces = append(namespaces, namespace.Name)
	}
	return namespaces
}

func getApiGroupsAndResources(loadedConfig models.Config) []models.GroupAndResourceNames {
	if len(loadedConfig.Resources) > 0 {
		logrus.Info("Found Specified Resources to scrape inside Config File, using those")
		return loadedConfig.Resources
	}
	logrus.Info("No SpecifiedResources Found in config, will scrape will resources")
	groupsAndResources := getServerPreferredNSResources()
	return groupsAndResources
}

func getServerPreferredNSResources() []models.GroupAndResourceNames {
	clientset := GetClient()
	var groupsAndResources []models.GroupAndResourceNames
	serverResources, err := clientset.ServerPreferredNamespacedResources()
	if err != nil {
		logrus.Errorf("Error listing groups present on the server: %s", err.Error())
	}
	for _, serverResource := range serverResources {
		groupDetails := strings.Split(serverResource.GroupVersion, "/")
		groupName := ""
		groupVersion := ""
		if len(groupDetails) == 1 {
			groupName = ""
			groupVersion = groupDetails[0]
		} else {
			groupName = groupDetails[0]
			groupVersion = groupDetails[1]
		}

		var resourceNames []string
		for _, apiResource := range serverResource.APIResources {
			resourceNames = append(resourceNames, apiResource.Name)
		}
		groupAndResource := models.GroupAndResourceNames{
			Group:         groupName,
			Version:       groupVersion,
			ResourceNames: resourceNames,
		}
		groupsAndResources = append(groupsAndResources, groupAndResource)
	}
	return groupsAndResources
}

func getNamespacedResourcesForAPIGroup(apiGroup string) []string {
	apiGroupsAndResources := getServerPreferredNSResources()

	for _, apiGroupInstance := range apiGroupsAndResources {
		if apiGroup == apiGroupInstance.Group {
			return apiGroupInstance.ResourceNames
		}
	}
	// no resources found return empty string
	var emptyList []string
	return emptyList
}
