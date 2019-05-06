package settings

import (
	"github.com/rs/zerolog/log"

	"github.com/spf13/viper"
)

func Setup(path string) {
	viper.SetConfigType("yaml")
	viper.SetConfigName("pompom")

	if path != "" {
		viper.AddConfigPath(path)
	}
	viper.AddConfigPath("/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./fixtures")

	viper.SetEnvPrefix("pompom")
	viper.AutomaticEnv()
	setDefaults()

	err := viper.ReadInConfig()
	if err != nil {
		log.Warn().Err(err).Msg("could not find a config file, using defaults")
	}
}

func setDefaults() {
	viper.SetDefault("output.pubsub.enabled", true)
	viper.SetDefault("output.pubsub.project", "")
	viper.SetDefault("output.pubsub.topic", "pompom")
	viper.SetDefault("output.gcs.enabled", false)
	viper.SetDefault("output.gcs.bucket", "")
	viper.SetDefault("output.gcs.prefix", "pompom/")
	viper.SetDefault("input.flatfile.enabled", false)
	viper.SetDefault("input.flatfile.location", "")
	viper.SetDefault("input.http.enabled", false)
	viper.SetDefault("input.grpc.enabled", false)
}

func setTests() {
	viper.Set("output.pubsub.enabled", true)
	viper.Set("output.pubsub.project", "test")
	viper.Set("output.pubsub.topic", "pompom")
	viper.Set("output.gcs.enabled", false)
	viper.Set("output.gcs.bucket", "test")
	viper.Set("output.gcs.prefix", "pompom/")
	viper.Set("input.flatfile.enabled", false)
	viper.Set("input.flatfile.location", "")
	viper.Set("input.http.enabled", false)
	viper.Set("input.grpc.enabled", false)
}
