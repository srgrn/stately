package main

import (
//"fmt"
//"github.com/BurntSushi/toml"
)

var cmdUpdate = &Command{
	UsageLine: "update [-f] [-p] [-r]",
	Short:     "updates all repositories in given project according to project definition file ",
	Long: `
The get command parse the project definition file and downloads the relevant sources into current directory.
Running without any flag will search for a file called projectdef.toml in the current directory.
The -p flag allows to specify a file to work on.
The -r flag allows to specify a remote file (from a git repository)
	`,
	CustomFlags: false,
}
