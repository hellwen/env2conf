package main

import (
	"fmt"
        "flag"
	"os"
	"log"
	"strings"
	"io/ioutil"
	"gopkg.in/ini.v1"
	"gopkg.in/yaml.v2"
)

var appName = "env2conf"

var configFile = flag.String("file", "", "point to the config file")
var fileType = flag.String("type", "ini", "config file type: ini, yaml")
var queryVersion = flag.Bool("version", false, "query version")

func toIni(configFile string) {
        cfg, err := ini.Load(configFile)
        if err != nil {
                log.Fatalf("%v Cant not load the config file: %v, error: %v", appName, configFile, err)
        }

	log.Printf("%v Init config file...", appName)
        for _, s := range cfg.Sections() {
                for _, k := range s.Keys() {
			var envName string
                        if strings.ToUpper(s.Name()) != "DEFAULT" {
                        	envName = strings.ToUpper(s.Name()) + "_" + strings.ToUpper(k.Name())
			} else {
                        	envName = strings.ToUpper(k.Name())
			}
			log.Printf("%v - %v = %v", appName, envName, k.String())

			envValue := os.Getenv(envName)
			if envValue != "" && envValue != k.String() {
				log.Printf("%v  => %v", appName, envValue)
				k.SetValue(envValue)
			}
                }
        }

	err = cfg.SaveTo(configFile)
        if err != nil {
                log.Fatalf("%v Save ini file error: %v", appName, err)
        }
	log.Printf("%v Finished init config file...", appName)
}

func toYaml(configFile string) {
	source, err := ioutil.ReadFile(configFile)
	if err != nil {
                log.Fatalf("error: %v", err)
	}

        m := make(map[interface{}]interface{})

        err = yaml.Unmarshal(source, &m)
        if err != nil {
                log.Fatalf("error: %v", err)
        }
        fmt.Printf("--- m:\n%v\n\n", m)

        d, err := yaml.Marshal(&m)
        if err != nil {
                log.Fatalf("error: %v", err)
        }
        fmt.Printf("--- m dump:\n%s\n\n", string(d))
}

func main() {
	flag.Parse()

	if *queryVersion {
		fmt.Printf("%v version 1.0\n", appName)
		return
	}

	log.Printf("%v version 1.0", appName)

	switch *fileType {
	case "ini":
		log.Printf("%v Ini file:%v", appName, *configFile)
		toIni(*configFile)
	case "yaml":
		log.Printf("%v Yaml file:%v", appName, *configFile)
		toYaml(*configFile)
	default:
		log.Printf("%v File type error:%v", appName, *fileType)
	}
}
