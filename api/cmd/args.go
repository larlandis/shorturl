package cmd

import "flag"

var (
	ConfigFile string
	Port       string
)

func ParseArgs() {
	flag.StringVar(&Port, "port", "8080", "application listen port")
	flag.StringVar(&ConfigFile, "config", "/cmd/config/config.yml", "config yaml file")
	flag.Parse()
}
