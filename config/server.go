package config

type Server struct {
	Port string `env:"PORT" envDefault:"8080"`
}
