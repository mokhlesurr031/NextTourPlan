package cmd

import (
	"github.com/NextTourPlan/internal/config"
	"github.com/spf13/cobra"
	"log"
)

var (
	// cfgFile store the configuration file name
	cfgFile                 string
	verbose, prettyPrintLog bool
	rootCmd                 = &cobra.Command{
		Use:   "tour",
		Short: "Tour Plan Backend Server",
		Long:  `Tour Plan Backend Server`,
	}
)

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	log.Println("Loading configurations")
	config.Init()
	log.Println("Configurations loaded successfully!")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
