package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName              string
	AppEnv               string
	SqlPrefix            string
	RedisAddr            string
	DBUserName           string
	DBPassword           string
	DBHostWriter         string
	DBHostReader         string
	DBPort               string
	DBName               string
	DBMaxOpenConnections int
	DBMaxIdleConnections int
	ServerPort           string
	EsURL                string
	EsPort               int
	SentryDSN            string
	SentrySamplingRate   float64
	AWSAccessKey         string
	AWSSecretKey         string
	JWT_SECRET           string
}

var config Config

// Should run at the very beginning, before any other package
// or code.
func init() {
	appEnv := os.Getenv("APP_ENV")
	if len(appEnv) == 0 {
		appEnv = "dev"
	}

	configFilePath := "./config/.env"

	switch appEnv {
	case "production":
		configFilePath = "./config/.env.prod"
		break
	case "stage":
		configFilePath = "./config/.env.stage"
		break
	}
	fmt.Println("reading env from: ", configFilePath)

	e := godotenv.Load(configFilePath)
	if e != nil {
		fmt.Println("error loading env: ", e)
		panic(e.Error())
	}
	config.AppName = os.Getenv("SERVICE_NAME")
	config.AppEnv = appEnv
	config.SqlPrefix = "/* " + config.AppName + " - " + config.AppEnv + "*/"
	config.RedisAddr = os.Getenv("REDIS_ADDR")
	config.DBUserName = os.Getenv("DB_USERNAME")
	config.DBHostReader = os.Getenv("DB_HOST_READER")
	config.DBHostWriter = os.Getenv("DB_HOST_WRITER")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	config.DBMaxIdleConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONENCTION"))
	config.DBMaxOpenConnections, _ = strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS"))
	config.ServerPort = os.Getenv("SERVER_PORT")
	config.EsURL = os.Getenv("ES_URL")
	config.EsPort, _ = strconv.Atoi(os.Getenv("ES_PORT"))
	config.SentryDSN = os.Getenv("SENTRY_DSN")
	config.SentrySamplingRate, _ = strconv.ParseFloat(os.Getenv("SENTRY_SAMPLING_RATE"), 64)
	config.AWSAccessKey = os.Getenv("AWS_KEY")
	config.AWSSecretKey = os.Getenv("AWS_SECRET")
	config.JWT_SECRET = os.Getenv("JWT_SECRET")
}

func Get() Config {
	return config
}

func IsProduction() bool {
	return config.AppEnv == "production"
}
