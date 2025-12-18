package config

import (
	"os"

	"github.com/naoina/toml"
)

// 서버 기본 설정
// 환경마다 다른 config, db 사용을 위해 분리

type Config struct {
	Server struct {
		Port string
	}
}

func NewConfig(filePath string) *Config {
	c := new(Config)
	//var c *Config{} -> 이런 식이 좋음
	if file, err := os.Open(filePath); err != nil { // filepath를 못받아오면
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil { // toml이 디코딩이 안되면
		panic(err)
	} else { // 다 되면, *Config 인스턴스를 제로값 구조체로 반환
		return c
	}
}
