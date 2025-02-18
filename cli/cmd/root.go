// Copyright ©, 2023-present, Lightspark Group, Inc. - All Rights Reserved
package cmd

import (
	"fmt"
	"os"

	"github.com/lightsparkdev/go-sdk/cli/cmd/offer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	viper.AutomaticEnv()
	viper.MustBindEnv("LIGHTSPARK_API_ENDPOINT")
	viper.MustBindEnv("LIGHTSPARK_API_TOKEN_CLIENT_ID")
	viper.MustBindEnv("LIGHTSPARK_API_TOKEN_CLIENT_SECRET")
	viper.MustBindEnv("LIGHTSPARK_NODE_ID")
	viper.MustBindEnv("LIGHTSPARK_MASTER_SEED_HEX")
	viper.MustBindEnv("LIGHTSPARK_BITCOIN_NETWORK")

	rootCmd.AddCommand(offer.OfferCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
