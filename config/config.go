package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/time/rate"
	"os"
	"time"
)

type SysConfig struct {
	ExternalApiKey         string        `env:"EXTERNAL_API_KEY"`
	LogMiddleware          bool          `env:"LOG_MIDDLEWARE"`
	HashCost               int           `env:"HASH_COST"`
	BruteForceLimit        int64         `env:"BRUTE_FORCE_LIMIT"`
	BruteForceMonthlyLimit int64         `env:"BRUTE_FORCE_MONTHLY_LIMIT"`
	BruteForceDuration     time.Duration `env:"BRUTE_FORCE_DURATION"`
	KeyBruteForceLimit     int64         `env:"KEY_BRUTE_FORCE_LIMIT"`
	KeyBruteForceDuration  time.Duration `env:"KEY_BRUTE_FORCE_DURATION"`
	LoginRateLimit         rate.Limit    `env:"LOGIN_RATE_LIMIT"`
}

var settings *SysConfig

// Get Accessor (Singleton)
func Get() *SysConfig {
	if settings == nil {
		setup()
	}
	return settings
}

// Set up the config
func setup() {
	// Rather than use a config file, let's set up some default here as we're lazy. These can be overwritten by the
	// env variables
	settings = &SysConfig{
		LogMiddleware:          true,
		HashCost:               10,
		BruteForceLimit:        5,
		BruteForceMonthlyLimit: 20,
		BruteForceDuration:     15 * time.Minute,
		KeyBruteForceLimit:     15,
		KeyBruteForceDuration:  15 * time.Minute,
		LoginRateLimit:         10,
	}

	err := cleanenv.ReadEnv(settings)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
