package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

var cmdGet = &Command{
	UsageLine: "get [-p] [-r]",
	Short:     "create the directory structure of a project according to project definition file",
	Long: `
The get command parse the project definition file and downloads the relevant sources into current directory.
Running without any flag will search for a file called projectdef.toml in the current directory.
The -p flag allows to specify a file to work on.
The -r flag allows to specify a remote file (from a git repository)
	`,
	CustomFlags: true,
}

var getP = cmdGet.Flag.String("p", "projectdef.toml", "Specify project definition file")
var getR = cmdGet.Flag.String("r", "", "remote url for project definition file")

func init() {
	cmdGet.Run = runGet // break init loop
}
func create(s source, ch chan bool) {
	fmt.Println("started on", s.Target)
	s.SourceType.create(s.Target, s.Url, s.Branch)
	fmt.Println("finished working on", s.Target)
	ch <- true
}

func runGet(cmd *Command, args []string) {
	var config Config
	if _, err := toml.DecodeFile(*getP, &config); err != nil {
		fmt.Println(err)
		return
	}
	done := make(chan bool, len(config.Sources))
	for _, source := range config.Sources {
		source.set_type()
		if source.Branch == "" {
			if config.BaseBranch == "" {
				source.Branch = source.SourceType.BaseBranchName
			} else {
				source.Branch = config.BaseBranch
			}
		}
		go create(source, done)

	}
	var wait = len(config.Sources)
	for wait > 0 {
		<-done
		wait -= 1
	}

}
