package main
	
import (
	"flag"
	"fmt"
	"github.com/srgrn/stately"
	"os"
	"github.com/BurntSushi/toml"
)

var versionFlag bool
//var configFile string

func main(){
	flag.BoolVar(&versionFlag,"version",false,"Show the version")
	flag.BoolVar(&versionFlag,"v",false,"Show the version")

	flag.Parse()
	if versionFlag {
		fmt.Println("Stately ",stately.Version())
		os.Exit(0)
	}	

	var config stately.Config
	if _, err := toml.DecodeFile("projectdef.toml", &config); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Name: %s\n",config.Name)
	config.Get()
	os.Exit(1)
}