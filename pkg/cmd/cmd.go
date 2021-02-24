package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tosone/logging"
	"github.com/unknwon/com"

	"github.com/tosone/golang-gin-template/pkg/database"
	"github.com/tosone/golang-gin-template/pkg/server"
	"github.com/tosone/golang-gin-template/pkg/version"
)

// RootCmd represents the base command when called without any sub commands
var RootCmd = &cobra.Command{
	Short: "golang-gin-template desc.",
	Long:  `golang-gin-template description.`,
}

const DefaultConfig = "/config/config.yml"

func init() {
	var config string
	RootCmd.PersistentFlags().StringVarP(&config, "config", "c", "./config.yml", "config file")

	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Restful API server for leetcode problems info.",
		Long:  `Restful API server for leetcode problems info.`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(_ *cobra.Command, _ []string) {
			var err error
			if err = initConfig(config); err != nil {
				logging.Error(err)
			}
			if err = database.Initialize(); err != nil {
				logging.Errorf("Got error: %+v", err)
				return
			}
			if err = server.Initialize(); err != nil {
				logging.Errorf("Got error: %+v", err)
			}
		},
	}
	RootCmd.AddCommand(serverCmd) // server commander

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Get version.",
		Long:  `The version that build detail information.`,
		Run: func(_ *cobra.Command, _ []string) {
			version.Initialize()
		},
	}
	RootCmd.AddCommand(versionCmd) // version commander

	RootCmd.Use = viper.GetString("AppName")
}

func initConfig(config string) (err error) {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	if com.IsFile(config) {
		viper.SetConfigFile(config)
	} else if config == "./config.yml" && com.IsFile(DefaultConfig) {
		viper.SetConfigFile(DefaultConfig)
	}
	if err := viper.ReadInConfig(); err != nil {
		logging.Error("Cannot find the specified config file.")
	}
	return
}
