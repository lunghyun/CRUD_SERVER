# CRUD_SERVER

Go 기반 REST API 서버로, 3계층 아키텍처 패턴을 구현한 학습 프로젝트입니다.

## 프로젝트 구조

```
CRUD_SERVER/
├── cmd/                    # 애플리케이션 진입점
│   └── main.go             # 메인 함수, --config 플래그 지원
│
├── internal/               # 내부 애플리케이션 코드
│   ├── cmd/                # 의존성 주입 및 생명주기 관리
│   │   └── cmd.go          # Graceful shutdown 포함
│   │
│   ├── config/             # 환경 설정 관리
│   │   ├── config.go       # .env 파일 로딩
│   │   ├── database.go     # DB 설정 구조체
│   │   ├── storage.go      # Storage 설정 구조체
│   │   └── server.go       # 서버 설정 구조체
│   │
│   ├── infra/              # 인프라 계층
│   │   └── db.go           # MySQL Connection Pool 관리
│   │
│   ├── network/            # HTTP 통신 계층 (Gin)
│   │   ├── root.go         # 라우터 및 Graceful Shutdown
│   │   ├── user.go         # 사용자 API 엔드포인트
│   │   └── utils.go        # HTTP 유틸리티 함수
│   │
│   ├── service/            # 비즈니스 로직 계층
│   │   ├── root.go         # 서비스 팩토리
│   │   └── user.go         # 사용자 비즈니스 로직
│   │
│   ├── repository/         # 데이터 접근 계층
│   │   ├── root.go         # 레포지토리 팩토리
│   │   ├── user_interface.go  # Repository 인터페이스
│   │   └── user_sql.go     # MySQL 구현체
│   │
│   └── types/              # 공유 타입 정의
│       ├── user.go         # 사용자 도메인 모델 및 DTO
│       ├── cerrors/        # 커스텀 에러 처리
│       │   └── errors.go
│       └── utils.go        # API 응답 구조체
│
├── migrations/             # 데이터베이스 마이그레이션
├── docs/                   # 프로젝트 문서
└── docker-compose.yml      # Docker 서비스 설정 (MySQL, MinIO)
```

## 아키텍처

**3계층 아키텍처 (3-Tier Architecture)**

```
main (cmd/)            ← 애플리케이션 진입점
    ↓
Cmd (internal/cmd/)    ← 의존성 주입 및 생명주기 관리
    ↓
HTTP Client
    ↓
Network Layer (Gin)    ← HTTP 요청/응답 처리, Graceful Shutdown
    ↓
Service Layer          ← 비즈니스 로직
    ↓
Repository Layer       ← 데이터 접근 추상화 (Interface)
    ↓
Infrastructure         ← DB Connection Pool 관리
    ↓
MySQL Database         ← 데이터 저장
```

## API 엔드포인트

모든 엔드포인트는 `/` 경로에서 HTTP 메서드로 구분:

| Method | 기능 | 요청 |
|--------|------|------|
| POST   | 사용자 생성 | `{"name": string, "age": int}` |
| GET    | 전체 사용자 조회 | - |
| PUT    | 사용자 수정 | `{"name": string, "age": int}` |
| DELETE | 사용자 삭제 | `{"name": string}` |

## 실행 방법

```bash
# 의존성 설치
go mod download

# MySQL 및 MinIO 시작 (Docker)
docker-compose up -d

# 서버 실행 (기본: .env)
go run cmd/main.go

# 또는 설정 파일 지정
go run cmd/main.go --config=/path/.env

# Graceful Shutdown: Ctrl+C
```

## 환경 변수

`.env` 파일에서 설정:

```env
# 서버
PORT=

# 데이터베이스 (MySQL)
DB_HOST=
DB_PORT=
DB_USER=
DB_PASSWORD=
DB_NAME=

# 스토리지 (MinIO)
STORAGE_API_PORT=
STORAGE_WEB_PORT=
STORAGE_USER=
STORAGE_PASSWORD=
```

## 주요 기술 스택

- **Language**: Go 1.25.0
- **Web Framework**: Gin (v1.11.0)
- **Database**: MySQL (Docker)
- **Database Driver**: go-sql-driver/mysql
- **Storage**: MinIO (Docker)
- **Config Management**: godotenv, caarlos0/env

## 디자인 패턴

- **3-Tier Architecture**: Network - Service - Repository 계층 분리
- **Repository Pattern**: 데이터 접근 추상화
- **Dependency Injection**: 계층 간 의존성 주입 (생성자 패턴)
- **Factory Pattern**: 각 계층별 팩토리 함수 (`NewXxx`)
- **Singleton Pattern**: Repository/Service/Network 레이어에서 `sync.Once` 사용
- **Graceful Shutdown**: Context 기반 Graceful Shutdown
  - 시그널 처리 (SIGINT, SIGTERM)
  - 진행 중인 요청 완료 대기 (5초 타임아웃)
  - DB 연결 정리