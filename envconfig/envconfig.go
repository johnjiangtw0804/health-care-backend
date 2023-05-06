package envconfig

import "github.com/kelseyhightower/envconfig"

type (
	Env struct {
		Port         int    `envconfig:"PORT" default:"5500" required:"true"`
		DATABASE_URL string `envconfig:"DATABASE_URL" required:"true"`
	}
)

func Process(env *Env) error {
	return envconfig.Process("", env)
}

func New() (*Env, error) {
	var env Env
	err := envconfig.Process("", &env)
	if err != nil {
		return nil, err
	}
	return &env, nil
}
