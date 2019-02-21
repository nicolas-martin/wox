package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

//Env ...
type Env struct {
}

//Load parses env variable and return an Env struct
func Load() Env {
	env := Env{}
	err := envconfig.Process("", &env)
	if err != nil {
		logrus.Panic(err.Error())
	}
	return env
}
