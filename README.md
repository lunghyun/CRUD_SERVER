# CRUD_SERVER

Go 기반 REST API 서버로, 3계층 아키텍처 패턴을 구현한 학습 프로젝트입니다.

## 프로젝트 구조

```
CRUD_SERVER/
├── init/               # 애플리케이션 초기화 및 진입점
│   ├── main.go        # 메인 함수, --config 플래그 지원
│   └── cmd/           # 의존성 주입 및 계층 초기화
│
├── config/            # 환경 설정 관리
│   ├── config.go      # .env 파일 로딩
│   └── server.go      # 서버 설정 구조체
│
├── network/           # HTTP 통신 계층 (Gin)
│   ├── root.go        # 라우터 설정 및 서버 시작
│   ├── user.go        # 사용자 API 엔드포인트
│   └── utils.go       # HTTP 유틸리티 함수
│
├── service/           # 비즈니스 로직 계층
│   ├── root.go        # 서비스 팩토리
│   └── user.go        # 사용자 비즈니스 로직
│
├── repository/        # 데이터 접근 계층
│   ├── root.go        # 레포지토리 팩토리
│   └── user.go        # 사용자 데이터 접근 (현재 인메모리)
│
└── types/             # 공유 타입 정의
    ├── user.go        # 사용자 도메인 모델 및 API 계약
    ├── cerrors/       # 커스텀 에러 처리
    └── utils.go       # API 응답 구조체
```

## 아키텍처

**3계층 아키텍처 (3-Tier Architecture)**

```
HTTP Client
    ↓
Network Layer (Gin)    ← HTTP 요청/응답 처리
    ↓
Service Layer          ← 비즈니스 로직
    ↓
Repository Layer       ← 데이터 접근
    ↓
Storage (In-Memory)    ← 데이터 저장 (향후 DB 연동 예정)
```

## API 엔드포인트

모든 엔드포인트는 `/` 경로에서 HTTP 메서드로 구분됩니다:

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

# 실행 (기본: ../.env)
cd init && go run main.go

# 또는 설정 파일 지정
cd init && go run main.go --config=/path/.env
```

## 환경 변수

`.env` 파일에서 설정:

```env
PORT=8080
```

## 주요 기술 스택

- **Web Framework**: Gin (v1.11.0)
- **Config Management**: godotenv, caarlos0/env
- **Language**: Go 1.25.0

## 디자인 패턴

- **Singleton Pattern**: 각 계층에서 `sync.Once` 사용
- **Factory Pattern**: 계층별 생성자 패턴
- **Dependency Injection**: 계층 간 의존성 주입
- **Repository Pattern**: 데이터 접근 추상화