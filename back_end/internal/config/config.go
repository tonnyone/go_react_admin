package config

type AppConfig struct {
	App struct {
		Name string
		Port int
		Mode string
	}
	DB struct {
		DSN string
	}
	Log struct {
		Level  string // info, debug, warn, error
		Format string // text, json
	}
}
