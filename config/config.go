package config

import (
	"fmt"
	"os"

	"github.com/naoina/toml"
)

// 서버 기본 설정
// 환경마다 다른 config, db 사용을 위해 분리

type Config struct {
	Server Server
}

func NewConfig(filePath string) (*Config, error) {
	c := new(Config)
	//var c *Config{} -> 이런 식이 좋음
	file, err := os.Open(filePath)
	if err != nil { // filepath를 못받아오면
		return nil, fmt.Errorf("config 파일(%s) 에러: %w", filePath, err)
	}

	defer func() { _ = file.Close() }()

	if err = toml.NewDecoder(file).Decode(c); err != nil { // toml이 디코딩이 안되면
		return nil, fmt.Errorf("디코딩 에러: %w", err)
	}
	// 다 되면, *Config 인스턴스를 제로값 구조체로 반환
	return c, nil
}
