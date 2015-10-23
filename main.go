// Stately is a commandline multiplatform tool for handling projects that contains several repositories.
// it will copy the functionalty given in the myrepos tool
package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/user"
)

var Config string

const VERSION = "v0.1.0"

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	var statelyCmd = &cobra.Command{
		Use:   "stately",
		Short: "Stately is a utility to assist with handling workspaces containing directories from several source repositories.",
		Long: `Stately was originaly designed to assist in getting multiple repositories at the same time.
		it has copied lots of functionalty from the go command and the repo tool, today it has added functionalty`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			var homedir string
			currentUser, err := user.Current()
			if err != nil {
				log.Warningln("cannot get current user will use . as homedir")
			} else {
				homedir = currentUser.HomeDir
				log.WithField("Home", homedir).Debugln("Got user homedir")

			}
			// create config if it is missing
			if Config == "" {
				viper.SetConfigName("stately")
				viper.AddConfigPath(".")
				viper.AddConfigPath("../")
				viper.AddConfigPath(homedir)
				log.WithField("Config", "stately").Debugln("searching for config file")
			} else {
				viper.SetConfigFile(Config)
				log.WithField("Config", Config).Debugln("Using config file from flag: ", Config)
			}
			err = viper.ReadInConfig() // Find and read the config file
			if err != nil {            // Handle errors reading the config file
				// _ = "breakpoint"
				log.WithField("err", err).Warningln("Error during reading config")
			}
		},
	}
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of stately",
		Long:  `This is the stately version number`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Stately Dev %s\n", VERSION)
		},
	}
	var registerCmd = &cobra.Command{
		Use:   "register",
		Short: "register a repository in stately config file",
		Long: `The register command allows you to add a new repo to the config file.
		it will first search a config file in the following order:
			current folder -> parent folder -> default config file.
		you can also specify the config file to work on specifically and it will be create if it doesn't exist`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Stately Dev %s\n", VERSION)
		},
	}
	statelyCmd.PersistentFlags().StringVarP(&Config, "config", "c", "", "Specific config file for stately.")
	statelyCmd.AddCommand(versionCmd)
	statelyCmd.AddCommand(registerCmd)
	statelyCmd.Execute()
}
