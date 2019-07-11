package main

import (
	"log"
	"os"

	"github.com/prologic/pastebin/client"

	"github.com/mitchellh/go-homedir"
	"github.com/namsral/flag"
)

const (
	defaultConfig     = "pastebin.conf"
	defaultUserConfig = "~/.pastebin.conf"
	defaultURL        = "http://localhost:8000"
)

func getDefaultConfig() string {
	path, err := homedir.Expand(defaultUserConfig)
	if err != nil {
		return defaultConfig
	}
	return path
}

func main() {
	var (
		config   string
		url      string
		insecure bool
	)

	flag.StringVar(&config, "config", getDefaultConfig(), "path to config")
	flag.StringVar(&url, "url", defaultURL, "pastebin service url")
	flag.BoolVar(&insecure, "insecure", false, "insecure (skip ssl verify)")

	flag.Parse()

	cli := client.NewClient(url, insecure)

	err := cli.Paste(os.Stdin)
	if err != nil {
		log.Printf("error posting paste: %s", err)
		os.Exit(1)
	}
}
