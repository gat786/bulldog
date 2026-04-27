package main

import (
	"context"
	"gat786/bulldog/config"
	"gat786/bulldog/exporter"
	"gat786/bulldog/kubernetes"
	logrus "gat786/bulldog/log"
	"os"

	// "gat786/bulldog/config"
	// "gat786/bulldog/exporter"
	// "gat786/bulldog/kubernetes"

	"github.com/joho/godotenv"

	"github.com/urfave/cli/v3"
)

func init() {
	// logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.DebugLevel)
	godotenv.Load()

	logrus.Info("Initialising Kubernetes Resource Scraper")
}

func scrape() {
	outputDir := exporter.GetOutputDir()
	logrus.Info("Using Output Directory: ", outputDir)
	logrus.Info("Scraping resources from the cluster")
	config := config.LoadConfig()
	kubernetes.GetResources(config, outputDir)
}

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			&cli.Command{
				Name: "scrape",
				Action: func(ctx context.Context, c *cli.Command) error {
					logrus.Info("Starting scraping")
					scrape()
					return nil
				},
			},
			&cli.Command{
				Name: "print-config",
				Action: func(ctx context.Context, c *cli.Command) error {
					logrus.Info("Printing Scrape config as available to bulldog")
					return nil
				},
			},
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		logrus.Fatal("Executing Bulldog failed, kindly check the command")
	}
}
