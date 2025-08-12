package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"sync"
)

const (
	DefaultConfigName = ""
)

var (
	configs sync.Map
	mx      sync.Mutex
)

func Get[T any](name string) *T {
	if val, ok := configs.Load(name); ok {
		return val.(*T)
	}

	mx.Lock()
	defer mx.Unlock()

	fn := "config/" + name + ".env"

	if err := godotenv.Load(fn); err != nil {
		log.Println("Error loading env file", err, fn)
	}

	cfg := new(T)

	if err := cleanenv.ReadEnv(cfg); err != nil {
		log.Fatalln("Error loading environment variables", err)
	}

	configs.Store(name, cfg)
	return cfg
}
