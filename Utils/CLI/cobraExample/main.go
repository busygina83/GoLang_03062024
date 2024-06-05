
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	var configPath string
	var name string
	rootCmd := cobra.Command{
		Use:     "rest-service-example",
		Version: "v1.0",
		Run: func(cmd *cobra.Command, args []string) {
			log.Infof("configPath == %s", configPath)
			log.Infof("name == %s", name)
		},
	}
	rootCmd.Flags().StringVarP(&configPath, "config", "c", "", "Config file path")
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Name of user")
	err := rootCmd.MarkFlagRequired("config")
	if err != nil {
		panic("rootCmd.MarkFlagRequired() failed")
	}
	err = rootCmd.MarkFlagRequired("name")
	if err != nil {
		panic("rootCmd.MarkFlagRequired() failed")
	}
	err = rootCmd.Execute()
	if err != nil {
		log.Info(err)
	}
}
