package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

// 서버 기본 설정
// 환경마다 다른 config, db 사용을 위해 분리

type Config struct {
	Server   Server
	Database Database
	Storage  Storage
}

func NewConfig(filePath string) (*Config, error) {
	c := new(Config)
	//var c *Config{} -> 이런 식이 좋음
	//file, err := os.Open(filePath)
	//if err != nil { // filepath를 못받아오면
	//	return nil, fmt.Errorf("config 파일(%s) 에러: %w", filePath, err)
	//}
	//defer func() { _ = file.Close() }()
	if err := godotenv.Load(filePath); err != nil { // .env 파일을 읽고 환경변수에 로드
		return nil, fmt.Errorf(".env를 가져오지 못했습니다: %w", err)
	}

	//if err = toml.NewDecoder(file).Decode(c); err != nil { // toml이 디코딩이 안되면
	//	return nil, fmt.Errorf("디코딩 에러: %w", err)
	//}
	if err := env.Parse(c); err != nil { // 환경 변수를 매핑
		return nil, fmt.Errorf("매핑 실패: %w", err)
	}

	return c, nil
}
