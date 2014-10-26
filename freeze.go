package main

import (
	"fmt"
	"os"
	//"github.com/BurntSushi/toml"
)

var cmdFreeze = &Command{
	UsageLine: "freeze targetfile",
	Short:     "freezes the directory structure of the current directory into a project definiton file ",
	Long: `
The freeze command allows you to freeze the current directory into a project definition file
it goes over each directory in the current directory tries to guess the relevant source control for the directioy,
it then creates prints the results into a file.
	`,
	CustomFlags: false,
	Run:         runFreeze,
}

func init() {

}

func runFreeze(cmd *Command, args []string) {
	if len(args) < 1 {
		fmt.Println("target file name missing")
		return
	}
	var config Config
	curr, err := os.Open(".")
	if err != nil {
		fmt.Println("Some Error")
	}

	folders, err := curr.Readdirnames(0)
	for _, f := range folders {
		fmt.Println(f)
		var s source
		sp := &s
		sp.set_type()
		config.Sources = append(config.Sources, s)
	}
}
