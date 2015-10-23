// Stately is a commandline multiplatform tool for handling projects that contains several repositories.
// it will copy the functionalty given in the myrepos tool
package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Config string

const VERSION = "v0.1.0"

func main() {
	var statelyCmd = &cobra.Command{
		Use:   "stately",
		Short: "Stately is a utility to assist with handling workspaces containing directories from several source repositories.",
		Long: `Stately was originaly designed to assist in getting multiple repositories at the same time.
		it has copied lots of functionalty from the go command and the repo tool, today it has added functionalty`,
		Run: func(cmd *cobra.Command, args []string) {
			// create config if it is missing
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
