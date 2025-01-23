package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	configFileBaseName = "app.conf"
)

type Config struct {
	App
	API
}

type App struct {
	Address string `yaml:"Address"`
}

type API struct {
	BaseUri          string `yaml:"Address"`
	TimeOutInSeconds int    `yaml:"TimeOutInSeconds"`
}

func NewConfig() (*Config, error) {
	InitLogger()

	v := viper.New()
	setViperConfig(v)

	err := v.ReadInConfig()
	if err != nil {
		zap.S().Errorf("Error reading config file: %v", zap.Error(err))
		return nil, err
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		zap.S().Errorf("Error unmarshalling config file: %v", zap.Error(err))
		return nil, err
	}
	return &c, nil
}

func setViperConfig(v *viper.Viper) {
	if len(os.Args) <= 1 {
		zap.L().Error("Missing environment argument. Usage: <binary> <environment>. Example: 'app dev'")
		return
	}

	env := strings.TrimSpace(os.Args[1])
	fileName := fmt.Sprintf("%s.%s", configFileBaseName, env)
	zap.S().Infof("Valsea API! Running app with config file: %s", fileName)

	workingDir, err := os.Getwd()
	if err != nil {
		zap.L().Error("Failed to get working directory", zap.Error(err))
	}

	v.SetConfigName(fileName)
	v.SetConfigType("yaml")
	v.AddConfigPath(filepath.Join(workingDir, "/config"))

	if err := v.ReadInConfig(); err != nil {
		zap.L().Error("Failed to read config file", zap.Error(err))
	}
}

func InitLogger() *zap.Logger {
	cores := BuildCores(zapcore.DebugLevel)
	rootLogger := zap.New(zapcore.NewTee(cores...))
	zap.RedirectStdLog(rootLogger)
	zap.ReplaceGlobals(rootLogger)
	return rootLogger
}

func BuildCores(minLevel zapcore.Level) []zapcore.Core {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000")

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		minLevel,
	)

	return []zapcore.Core{
		core,
	}
}
