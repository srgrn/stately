package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

var cmdFreeze = &Command{
	UsageLine: "freeze [-b baseBranch] <targetfile>",
	Short:     "freezes the directory structure of the current directory into a project definiton file ",
	Long: `
The freeze command allows you to freeze the current directory into a project definition file
it goes over each directory in the current directory tries to guess the relevant source control for the directioy,
it then creates prints the results into a file.
	`,
	CustomFlags: true,
}

var freezeB = cmdFreeze.Flag.String("b", "", "Specify base branch if required")

func init() {
	cmdFreeze.Run = runFreeze // break init loop
}

func runFreeze(cmd *Command, args []string) {
	if len(args) < 1 {
		fmt.Println("target file name missing")
		return
	}
	newFileName := args[0]
	var config Config
	curr, err := os.Open(".")
	if err != nil {
		fmt.Println("Some Error")
	}

	files, err := curr.Readdir(0)
	for _, f := range files {
		if f.IsDir() {
			//fmt.Println(f.Name())
			var getterS, s source
			getterS.Target = f.Name()
			fmt.Println("working on", f.Name())
			err = getterS.set_type()
			if err == nil {
				s.Target = getterS.Target
				s.Url = getterS.SourceType.urlGetFunc(getterS.Target, getterS.SourceType)
				s.Branch = getterS.SourceType.branchGetFunc(getterS.Target, getterS.SourceType)
				//s.Branch = sp.get_branch()
				config.Sources = append(config.Sources, s)
			} else {
				fmt.Println(err)
			}

		}
	}
	pwd, err := os.Getwd()
	if err != nil {
	}
	config.Name = filepath.Base(pwd)
	if *freezeB != "" {
		config.BaseBranch = *freezeB
	}
	file, err := os.Create(newFileName)
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	encoder := toml.NewEncoder(file)
	encoder.Encode(config)

}
