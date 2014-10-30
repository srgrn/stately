package main

import (
	"fmt"
)

var cmdGet = &Command{
	UsageLine: "get <path to project definition file>",
	Short:     "create the directory structure of a project according to project definition file",
	Long: `
The get command parse the project definition file and downloads the relevant sources into current directory.
Running without any flag will search for a file called projectdef.toml in the current directory.
	`,
	CustomFlags: false,
}

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

	config, err := proccessConfig(args[0]) // moved to proccessConfig.go to support additional types and keywords
	if err != nil {
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
