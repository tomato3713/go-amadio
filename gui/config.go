package gui

type DebugConfig struct {
	Enable   bool   `yaml:"enable"`
	FilePath string `yaml:"file"`
	Level    int    `yaml:"level"`
}

type DisplayConfig struct {
	Timezone string `yaml:"timezone"`
}

type Config struct {
	ConfigDir  string
	ConfigFile string
	Debug      DebugConfig   `yaml:"debug"`
	Display    DisplayConfig `yaml:"display"`
	OpenCmd    string        `yaml:"open_cmd"`
}

func DefaultConfig() Config {
	return Config{
		Debug: DebugConfig{
			Enable:   false,
			FilePath: "amadio.log",
			Level:    0,
		},
		OpenCmd: "open",
	}
}
