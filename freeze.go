package main

import (
//"fmt"
//"github.com/BurntSushi/toml"
)

var cmdFreeze = &Command{
	UsageLine: "freeze targetfile",
	Short:     "freezes the directory structure of the current directory into a project definiton file ",
	Long: `
The freeze command allows you to freeze the current directory 
	`,
	CustomFlags: false,
}
