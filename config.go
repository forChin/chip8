package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type config struct {
	WindowW          int32 `env:"WINDOW_WIDTH" envDefault:"320"`
	WindowH          int32 `env:"WINDOW_HEIGHT" envDefault:"640"`
	OpcodesPerSecond int   `env:"OPCODES_PER_SECOND" envDefault:"300"`

	ROMPath string `env:"GAME_ROM_PATH"`
}

func newConfig() (*config, error) {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", cfg)
	return &cfg, nil
}
