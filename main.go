package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/tomato3713/amadio/gui"
	"gopkg.in/yaml.v2"
)

var (
	debugEnable    = flag.Bool("debug", false, "enable debug mode.")
	debugLevel     = flag.Int("debug_level", 0, "debug level.") // 0: Info, 1: Normal, 2: All
	configFilePath = flag.String("config", "", "config file")
)

func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func initConfig() (gui.Config, error) {
	var config gui.Config
	config = gui.DefaultConfig()

	// if specified config file path, use this path.
	// if not, use ${OS standard config dir}/amadio/amadio.yaml.
	if len(*configFilePath) == 0 {
		configDir, err := os.UserConfigDir()
		if err != nil {
			return config, err
		}
		config.ConfigDir = filepath.Join(configDir, "amadio")
		config.ConfigFile = filepath.Join(config.ConfigDir, "amadio.yaml")
	} else {
		config.ConfigDir = filepath.Dir(*configFilePath)
		config.ConfigFile = *configFilePath
	}

	if IsExist(config.ConfigFile) {
		b, err := ioutil.ReadFile(config.ConfigFile)
		if err != nil {
			return config, err
		}
		if err := yaml.Unmarshal(b, &config); err != nil {
			return config, err
		}
	}

	return config, nil
}

func run() int {
	flag.Parse()

	config, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", config)

	if err := gui.New(config).Run(); err != nil {
		log.Fatal(err)
	}

	return 0
}

func main() {
	os.Exit(run())
}
