package fagblog

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	// The port on which the server will listen
	Port int

	// The directory where the templates are stored
	TemplateDir string

	// The directory where the blog content is stored
	ContentDir string

	StaticDir string
}

func DefaultConfig() Config {
	return Config{
		Port:        8000,
		TemplateDir: "/usr/local/share/fagblog/templates",
		StaticDir:   "/usr/local/share/fagblog/static",
		ContentDir:  "/var/lib/fagblog",
	}
}

func ConfigFromToml(path string) (Config, error) {
	config := DefaultConfig()

	_, err := toml.DecodeFile(path, &config)

	if err != nil {
		log.Printf("Error decoding TOML file: %v", err)
		return config, err
	}

	return config, nil
}

func LoadConfig() (Config, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Printf("Error getting user config dir: %v", err)
		log.Println("Using /etc/")
		userConfigDir = "/etc"
	}

	configSearchPaths := []string{
		"config.toml",
		userConfigDir + "/fagblog/config.toml",
		"/etc/fagblog/config.toml",
	}

	for _, p := range configSearchPaths {
		if _, err := os.Stat(p); err == nil {
			return ConfigFromToml(p)
		}
	}

	log.Println("No config file found, using default config")
	return DefaultConfig(), nil
}
