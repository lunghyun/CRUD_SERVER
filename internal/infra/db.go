package infra

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lunghyun/CRUD_SERVER/internal/config"
)

type DB struct {
	Conn *sql.DB
}

func NewDB(cfg config.Database) (*DB, error) {
	// 1. dsn 설정
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
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

	// 5. 성공 시에만 인스턴스에 할당
	log.Printf("DB 연결 성공: %s@%s:%s/%s\n", cfg.User, cfg.Host, cfg.Port, cfg.Name)

	return &DB{Conn: conn}, nil
}

func (db *DB) Close() error {
	return db.Conn.Close()
}
