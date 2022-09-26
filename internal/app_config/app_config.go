package appconfig

import (
	"micro_service_phone/pkg/file_logger"

	"github.com/spf13/viper"
)

func InitConfig(logger *file_logger.FileLogger) error {
	logger.Traceln("Start config initialization")
	
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}