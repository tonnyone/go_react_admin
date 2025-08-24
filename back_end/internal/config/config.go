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
}
