package main

import (
	"crypto/tls"
	"fmt"
	"github.com/BurntSushi/toml"
	"net/http"
	"strings"
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

func proccessConfig(path string) (Config, error) {
	var config Config
	if strings.HasPrefix(strings.ToLower(path), "http") {
		fmt.Println("working with remote definition file")
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}
		res, _ := client.Get(path)

		_, err := toml.DecodeReader(res.Body, &config)
		if err != nil {
			fmt.Println(err)
			return config, err
		}
		res.Body.Close()
		return config, nil
	}
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		fmt.Println(err)
		return config, err
	}
	return config, nil
}

func runGet(cmd *Command, args []string) {

	config, err := proccessConfig(args[0])
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
