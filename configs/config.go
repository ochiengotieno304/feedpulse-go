package configs

import (
	"log"
	"os"
	"regexp"

	"github.com/spf13/viper"
)

type Configs struct {
	DatabaseUrl        string `mapstructure:"DATABASE_URL"`
	RapidApiProySecret string `mapstructure:"RAPID_API_PROXY_SECRET"`
}

const PROJECT_DIR = "feedpulse-go"

func LoadConfig() (config Configs, err error) {

	databaseUrl, rapidApiProxySecret := os.Getenv("DATABASE_URL"), os.Getenv("RAPID_API_PROXY_SECRET")

	if databaseUrl != "" && rapidApiProxySecret != "" {
		config.DatabaseUrl = databaseUrl
		config.RapidApiProySecret = rapidApiProxySecret
		return config, nil
	}

	projectName := regexp.MustCompile(`^(.*` + PROJECT_DIR + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	viper.AddConfigPath(string(rootPath))
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	err = viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return config, nil
}
