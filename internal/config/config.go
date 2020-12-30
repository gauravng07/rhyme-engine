package config

import "github.com/spf13/viper"

const (
	Port     = "PORT"
	Env      = "ENV"
	LogLevel = "LOG_LEVEL"
	BaseURL  = "BASE_URL"
	RefFile	 = "REF_File"
)

func init() {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetDefault(Port, "8080")
	viper.SetDefault(Env, "dev")
	viper.SetDefault(LogLevel, "debug")
	viper.SetDefault(BaseURL, "/api/v1")
	viper.SetDefault(RefFile, "assets/ref_rhyme_word.json")
}

func ReadConfig(env string) error {
	viper.SetConfigFile("app-" + env + ".yaml")
	return viper.ReadInConfig()
}
