package config

import (
	"os"

	"github.com/naoina/toml"
)

// 서버 기본 설정
// 환경마다 다른 config, db 사용을 위해 분리

type Config struct {
	Server struct {
		Port int64
	}
}

func NewConfig(filePath string) *Config {
	c := new(Config)
	//var c *Config{} -> 이런 식이 좋음
	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
