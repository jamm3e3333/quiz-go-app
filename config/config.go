package config

import "github.com/caarlos0/env"

type GRPCServer struct {
	Port                int64 `env:"GRPC_PORT" envDefault:"8088"`
	ShouldUseReflection bool  `env:"GRPC_USE_REFLECTION" envDefault:"true"`
}

type Logger struct {
	Level              string `env:"LOG_LEVEL" envDefault:"info"`
	ShouldUseDevelMode bool   `env:"LOG_USE_DEVEL_MODE" envDefault:"true"`
}

type Config struct {
	GRPCServer *GRPCServer
	Logger     *Logger
}

func NewParseConfigForENV() (*Config, error) {
	cfg := &Config{
		GRPCServer: &GRPCServer{},
		Logger:     &Logger{},
	}

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
