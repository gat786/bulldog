package main

import (
	logrus "gat786/bulldog/log"

	"gat786/bulldog/config"
	"gat786/bulldog/exporter"
	"gat786/bulldog/kubernetes"

	"github.com/joho/godotenv"
)

func init() {
	// logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	godotenv.Load()

	logrus.Info("Initialising Kubernetes Resource Scraper")
}

func main() {
	outputDir := exporter.GetOutputDir()
	logrus.Info("Using Output Directory: ", outputDir)
	logrus.Info("Scraping resources from the cluster")
	config := config.LoadConfig()
	kubernetes.GetResources(config, outputDir)
}
