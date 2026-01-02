package config

type Storage struct {
	StorageAPIPort  string `env:"STORAGE_API_PORT"`
	StorageWebPort  string `env:"STORAGE_WEB_PORT"`
	StorageUser     string `env:"STORAGE_USER"`
	StoragePassword string `env:"STORAGE_PASSWORD"`
}
