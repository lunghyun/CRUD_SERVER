package config

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
}

var (
	dbMutex      sync.Mutex
	dbConnection *sql.DB
)

func (db *Database) NewConnection() (*sql.DB, error) {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	// 이미 연결 성공 -> 기존 커넥션 재사용
	if dbConnection != nil {
		return dbConnection, nil
	}

	// 1. dsn 설정
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)

	// 2. db connection 생성
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("DB 연결 실패: %w", err)
	}

	// 3. db connection pool 설정 (Spring HikariCP 기본값 적용)
	conn.SetMaxOpenConns(10)                  // maximum-pool-size
	conn.SetMaxIdleConns(10)                  // minimum-idle
	conn.SetConnMaxLifetime(30 * time.Minute) // max-lifetime
	conn.SetConnMaxIdleTime(10 * time.Minute) // idle-timeout

	// 4. connection 검증 (실제 db와 연결되었는지)
	if err = conn.Ping(); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("DB ping 실패: %w", err)
	}

	// 5. 성공 시에만 전역 변수 할당
	dbConnection = conn
	fmt.Printf("DB 연결 성공: %s@%s:%s/%s\n", db.User, db.Host, db.Port, db.Name)

	return dbConnection, nil
}

func (db *Database) Close() error {
	if dbConnection != nil {
		return dbConnection.Close()
	}
	return nil
}
