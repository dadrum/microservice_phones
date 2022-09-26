package environment

import (
	appconfig "micro_service_phone/internal/app_config"
	argsparser "micro_service_phone/internal/args_parser"
	"micro_service_phone/internal/countries"
	apilayercountries "micro_service_phone/internal/countries/apilayer_countries"
	phonevalidator "micro_service_phone/internal/phone_validator"
	apilayervalidator "micro_service_phone/internal/phone_validator/apilayer_validator"
	sendsms "micro_service_phone/internal/send_sms"
	twilliosendsms "micro_service_phone/internal/send_sms/twilio_send_sms"
	cacherepo "micro_service_phone/pkg/cache_repo"
	freecache "micro_service_phone/pkg/cache_repo/free_cache"
	"micro_service_phone/pkg/file_logger"

	"github.com/spf13/viper"
)

type Environment struct {
	Cache          *cacherepo.ICacheRepository
	Logger         *file_logger.FileLogger
	Countries      *countries.ICountriesRepository
	Phonevalidator *phonevalidator.IPhoneValidator
	SendSms        *sendsms.ISendSms
}

// --------------------------------------------------------------------------------------
// constructor for handler`s dependencies
func newEnvironment(
	cache *cacherepo.ICacheRepository,
	logger *file_logger.FileLogger,
	countries *countries.ICountriesRepository,
	phonevalidator *phonevalidator.IPhoneValidator,
	sendSms *sendsms.ISendSms,
) *Environment {
	return &Environment{
		Cache:          cache,
		Logger:         logger,
		Countries:      countries,
		Phonevalidator: phonevalidator,
		SendSms:        sendSms,
	}
}

// --------------------------------------------------------------------------------------
func InitEnvironment() *Environment {

	// initialize command-line arguments
	argsparser.Init()

	// initialize new logger entity with selected level
	logger, _ := file_logger.Init(argsparser.LogLevelValue)

	// initialize app config
	initConfigErr := appconfig.InitConfig(logger)
	if initConfigErr != nil {
		logger.Fatalln("Error on config initialization with: ", initConfigErr.Error())
	} else {
		logger.Debugln("Config is initialized")
	}

	// initialize cache repository with specified expire time and cache size
	// *cacherepo.ICacheRepository
	cacheRepo := freecache.InitCacheRepository(
		*argsparser.CacheExpireTime,
		*argsparser.Cachesize,
		logger)

	// initialize countries repository
	// *countries.ICountriesRepository
	countriesRepo := apilayercountries.NewApiLayerCountries(
		viper.GetString("countries.apiKey"),
		logger,
	)

	// service to validate phone number
	phonevalidator := apilayervalidator.NewApiLayerPhoneValidator(
		viper.GetString("phoneValidator.apiKey"),
		logger,
		&cacheRepo,
	)

	// service to send sms
	smsSendService := twilliosendsms.NewTwilioSendSms(
		viper.GetString("sendSms.accountSid"),
		viper.GetString("sendSms.authToken"),
		viper.GetString("sendSms.fromPhone"),
		logger,
		&cacheRepo,
	)

	logger.Debugln("Environment is initialized")
	logger.Debugln("---------------------------------------")
	return newEnvironment(
		&cacheRepo,
		logger,
		&countriesRepo,
		&phonevalidator,
		&smsSendService,
	)
}
