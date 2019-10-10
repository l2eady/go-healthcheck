package app

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type (
	Configs struct {
		Stage         Stage
		ReportService ReportService `mapstructure:"heathycheckreport_service"`
	}

	ReportService struct {
		Address        string `mapstructure:"address"`
		ReportEndPoint string `mapstructure:"report_endpoint"`
		AccessToken    string `mapstructure:"access_token"`
	}
)
type Stage string

func (s Stage) String() string {
	return string(s)
}

const (
	StageLocal Stage = "local"
	StageDEV   Stage = "dev"
	StageSIT   Stage = "sit"
	StageProd  Stage = "prod"
)

func ParseStage(s string) Stage {
	switch s {
	case "local", "localhost", "l":
		return StageLocal
	case "dev", "develop", "development", "d":
		return StageDEV
	case "sit", "staging", "s":
		return StageSIT
	case "prod", "production", "p":
		return StageProd
	}
	return StageLocal
}

func New(path, state string) (*Configs, error) {

	var conf *Configs
	stage := ParseStage(state)
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(stage.String())
	if err := v.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to read config")
	}

	if err := v.Unmarshal(&conf); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal config")
	}
	conf.Stage = stage
	return conf, nil

}
