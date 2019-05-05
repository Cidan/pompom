package settings

import (
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

func Setup(path string) {
	viper.SetConfigFile("pompom")
	viper.SetConfigType("yaml")

	if path != "" {
		viper.AddConfigPath(path)
	}
	viper.AddConfigPath("/")
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("pompom")
	viper.AutomaticEnv()
	setDefaults()

	err := viper.ReadInConfig()
	if err != nil {
		log.Warn().Err(err).Msg("could not find a config file, using defaults")
	}
}

func setDefaults() {
	viper.SetDefault("pubsub.project", "")
	viper.SetDefault("pubsub.topic", "pompom")
	viper.SetDefault("gcs.bucket", "")
	viper.SetDefault("gcs.prefix", "pompom/")
	viper.SetDefault("input.flatfile.enabled", false)
	viper.SetDefault("input.flatfile.location", "")
	viper.SetDefault("input.http.enabled", false)
	viper.SetDefault("input.grpc.enabled", false)
}
