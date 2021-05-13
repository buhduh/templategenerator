/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	logger "github.com/buhduh/go-logger"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"templategenerator/constants"
	"templategenerator/document"
	"templategenerator/server"
)

var cfgFile, templateDir, logLevel, logLocation, port string
var showVersion bool

var mainLogger logger.Logger

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "templategenerator",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Printf(
				"%s version %s",
				constants.APPLICATION,
				constants.VERSION,
			)
			os.Exit(0)
		}
		mainLogger.Debugf("starting %s version: %s", constants.APPLICATION, constants.VERSION)
		docHelper := document.NewDocumentHelper(mainLogger)
		templateFS := os.DirFS(templateDir)
		myServer := server.NewServer(port, templateFS, mainLogger, docHelper)
		mainLogger.Debugf("starting server on port %s", port)
		if err := myServer.Run(); err != nil {
			mainLogger.Fatalf("recieved fatal error, '%s', quitting", err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	initFlags()
	initLogging()
}

func initFlags() {
	rootCmd.PersistentFlags().StringVar(
		&logLocation, "log-location", constants.DEF_LOG_LOCATION,
		fmt.Sprintf(
			"where %s will write logs",
			constants.APPLICATION,
		),
	)
	rootCmd.PersistentFlags().StringVar(
		&logLevel, "log-level", constants.LOG_LEVEL,
		fmt.Sprintf(
			"the log level of %s",
			constants.APPLICATION,
		),
	)
	rootCmd.PersistentFlags().StringVarP(
		&port, "port", "p", constants.DEF_PORT,
		"port to bind localhost to",
	)
	rootCmd.PersistentFlags().StringVar(
		&cfgFile, "config", "", "config file (default is $HOME/.templategenerator.yaml)",
	)
	rootCmd.PersistentFlags().StringVarP(
		&templateDir, "template", "t", constants.TEMPLATE_DIR,
		fmt.Sprintf(
			"where %s looks for the main templates directory",
			constants.APPLICATION,
		),
	)
	rootCmd.PersistentFlags().BoolVar(
		&showVersion, "version", false,
		fmt.Sprintf(
			"prints the version and exits",
		),
	)
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initLogging() {
	var logWriter = os.Stdout
	var err error
	if logLocation != constants.STD_OUT {
		logWriter, err = os.Open(logLocation)
		if err != nil {
			fmt.Fprintf(
				os.Stderr,
				"could not open %s for logging, using %s instead",
				logLocation, constants.STD_OUT,
			)
			logWriter = os.Stdout
		}
	}
	lLevel, err := logger.NewLogLevel(logLevel)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed determining log level, setting to %s\n", logger.WARN)
		lLevel = logger.WARN
	}
	mainLogger = logger.NewLogger(lLevel, "main logger", logWriter)
	mainLogger.Tracef("initialized logger")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".templategenerator" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".templategenerator")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
