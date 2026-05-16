package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App     AppConfig
		HTTP    HTTPConfig
		Log     LogConfig
		PG      PGConfig
		GRPC    GRPCConfig
		RMQ     RMQConfig
		NATS    NATSConfig
		JWT     JWTConfig
		Metrics MetricsConfig
		Swagger SwaggerConfig
	}

	AppConfig struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	HTTPConfig struct {
		Port           string `env:"HTTP_PORT,required"`
		UsePreforkMode bool   `env:"HTTP_USE_PREFORK_MODE" envDefault:"false"`
	}

	LogConfig struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	PGConfig struct {
		PoolMax int    `env:"PG_POOL_MAX,required"`
		URL     string `env:"PG_URL,required"`
	}

	GRPCConfig struct {
		Port string `env:"GRPC_PORT,required"`
	}

	RMQConfig struct {
		ServerExchange string `env:"RMQ_RPC_SERVER,required"`
		ClientExchange string `env:"RMQ_RPC_CLIENT,required"`
		URL            string `env:"RMQ_URL,required"`
	}

	NATSConfig struct {
		ServerExchange string `env:"NATS_RPC_SERVER,required"`
		URL            string `env:"NATS_URL,required"`
	}

	JWTConfig struct {
		Secret      string        `env:"JWT_SECRET,required"`
		TokenExpiry time.Duration `env:"JWT_TOKEN_EXPIRY" envDefault:"24h"`
	}

	MetricsConfig struct {
		Enabled bool `env:"METRICS_ENABLED" envDefault:"true"`
	}

	SwaggerConfig struct {
		Enabled bool `env:"SWAGGER_ENABLED" envDefault:"false"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
