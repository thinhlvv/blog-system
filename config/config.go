package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config represents the configurations for the application.
type Config struct {
	Env    string `envconfig:"ENV" default:"development"`
	Server struct {
		Port         string `envconfig:"SERVER_PORT" default:":8080"`
		ReadTimeout  int    `envconfig:"SERVER_READ_TIMEOUT" default:"300"`
		WriteTimeout int    `envconfig:"SERVER_WRITE_TIMEOUT" default:"300"`
	}
	Mysql struct {
		Name     string `envconfig:"DB_NAME" required:"true"`
		User     string `envconfig:"DB_USER" required:"true"`
		Password string `envconfig:"DB_PASS" required:"true"`
	}
	// Version string    `envconfig:"VERSION" required:"true"`
	Uptime time.Time `ignored:"true"`
	// BuildDate time.Time `envconfig:"BUILD_DATE" default:"2006-01-02T15:04:05Z"`
}

// New initialises a new configuration.
func New() *Config {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	// Setup the uptime.
	cfg.Uptime = time.Now()
	return &cfg
}
