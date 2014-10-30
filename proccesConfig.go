package main

import (
	"crypto/tls"
	"fmt"
	"github.com/BurntSushi/toml"
	"net/http"
	"strings"
)

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
