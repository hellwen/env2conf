package main

import (
	"fmt"
        "flag"
	// "gopkg.in/ini.v1"
	// "gopkg.in/yaml.v2"
)

var configFile = flag.String("file", "", "point to the config file")
var fileType = flag.String("type", "ini", "config file type: ini, yaml")

func main() {
	flag.Parse()

	switch fileType {
	case "ini":
		fmt.Println("ini file:", *configFile)
	case "yaml":
		fmt.Println("yaml file:", *configFile)
	default:
		fmt.Println("file type error:", *fileType)
	}
}
