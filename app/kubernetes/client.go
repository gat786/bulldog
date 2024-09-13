package kubernetes

import (
	"os"
	"path/filepath"

	logrus "gat786/bulldog/log"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func getKubeConfig() *rest.Config {
	// Initialize Kubernetes client
	k8sServiceHost := os.Getenv("KUBERNETES_SERVICE_HOST")
	var kubeconfig *rest.Config

	if k8sServiceHost != "" {
		logrus.Debug("Loading Credentials from inside the cluster")
		inClusterConfig, err := rest.InClusterConfig()
		if err != nil {
			logrus.Fatalf("Error loading config from incluster elements")
			os.Exit(1)
		}
		kubeconfig = inClusterConfig
	} else {
		logrus.Debug("Loading Credentials from $home/.kube directory")
		configPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
		pathConfig, err := clientcmd.BuildConfigFromFlags("", configPath)

		if err != nil {
			logrus.Fatalf("Error building kubeconfig: %v", err)
			os.Exit(1)
		}
		kubeconfig = pathConfig
	}
	return kubeconfig
}

func GetClient() *kubernetes.Clientset {
	kubeconfig := getKubeConfig()
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		logrus.Fatalf("Error creating Kubernetes client: %v", err)
	}
	return clientset
}

func GetDynamicClient() *dynamic.DynamicClient {
	kubeconfig := getKubeConfig()
	dynamicClient, err := dynamic.NewForConfig(kubeconfig)
	if err != nil {
		logrus.Fatalf("Error creating dynamic client: %v", err)
	}
	return dynamicClient
}
