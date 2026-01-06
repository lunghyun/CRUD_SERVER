# TODO
## 고려한 문제점
### 1. Context 인자 추가
- graceful shutdown -> 현재는 http 서버 레벨만
- 따라서 db 작업에도 shutdown 신호를 전파
- 트랜잭션 구현을 위해

### 2. Transaction
- ACID 문제 해결 가능

### 3. `sqlc` 도입
- query가 string으로 저장됨(타입 안정성) 
- 안정성을 위해 orm 도입 고려 (`ent` vs `sqlc`)
- go스러운 구성을 위해 최소한의 프레임워크 도입 (`sqlc`)

## 그래서 뭘 할까
### 1. Context 미들웨어 구성
- [x] ***is it Done?***

### 2. Transaction 구현
- [ ] ***is it Done?***
#### 트랜잭션 로직은 2가지 케이스으로 이루어짐
1. 정상 케이스
    - tx 시작(tx, err := db.Begin())
    - 작업 1..
    - 작업 2..
    - 작업 3..
    - 커밋 (tx.Commit())
2. 에러 케이스
   - tx 시작
   - 작업 1..
   - 작업 2..
   - 작업 3 실패
   - 롤백(tx.Rollback())
> #### Rollback을 defer로 지정해도 될까?
> 된다고 생각함.
> 1. 조기 리턴할 경우, Rollback 
>    - 의도한 실행
> 2. 정상 리턴할 경우, Commit후 Rollback
>    - Commit or Rollback이 한 번이라도 성공 -> 해당 트랜잭션은 종료됨
>    - 의도한 실행

알아야할 메서드
```go
tx, err := db.BeginTx(ctx, nil)
```
- 트랜잭션 시작
- *sql.Tx 객체 반환
- 커넥션을 하나 잡고, 트랜잭션 작업으로 기록
```go
tx.ExecContext(ctx, query, ...)
```
- 쿼리 실행
- DB 반영 x, 임시 저장
- db.ExecContext와 동일하지만 tx 범위 안에서 실행됨
```go
tx.Commit()
```
- 커밋
- 임시 저장한 db 작업 반영
```go
tx.Rollback()
```
- 롤백
- 모든 임시 작업 취소: db 상태를 트랜잭션 시작 전으로
> 컨텍스트 인자 값은, 해당 트랜잭션이 ctx에 종속된다는 뜻

#### 구현은 인터페이스로(transaction, 일반 db용 메서드가 나눠짐)

### 3. `sqlc` 도입
- [ ] ***is it Done?***
- ORM은 아님
  - sql 추상화
- sql기반 코드 생성기
  - sql 직접 사용
  - 런타임 오버헤드 없음

#### 사용법 ([SQLC Docs](https://docs.sqlc.dev/en/stable/tutorials/getting-started-mysql.html))
1. 환경에 설치
    ```go
    go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
    ```
2. `sqlc.[yaml | yml]` or `sqlc.json`작성
    ```yaml
    version: 2
    sql:
      - engine: "mysql" # 사용할 driver
        queries: "internal/queries" # 사용할 쿼리문 -> 디렉터리도 가능
        schema: "migrations" # 테이블이나 기본적인 스키마 구성 -> 디렉터리도 가능
        gen:
          go:
            package: "sqlc" # 만들 패키지 이름
            out: "internal/sqlc" # 만들 디렉터리 이름
            emit_json_tags: true # JSON 태그 추가
            emit_interface: true # Repository 인터페이스 자동 생성
            emit_empty_slices: true # nil 대신 빈 슬라이스 반환
    ```
3. Queries에 사용할 쿼리 작성(internal/queries/user.go)
    ```sql
    -- name: CreateUser :exec
    INSERT INTO users (name, age) VALUES (?, ?);
    
    -- name: GetAllUsers :many
    SELECT id, name, age FROM users;
    
    -- name: GetUserByName :one
    SELECT id, name, age FROM users WHERE name = ? LIMIT 1;
    
    -- name: UpdateUserAge :execresult
    UPDATE users SET age = ? WHERE name = ?;
    
    -- name: DeleteUserByName :execresult
    DELETE FROM users WHERE name = ?;
    
    ```
   - `-- name: 함수명 :리턴타입`: sqlc가 인식하는 특수 주석
   - `:exec`: 결과 없음 (INSERT 등)
   - `:execresult`: sql.Result 반환 (RowsAffected 체크 가능)
   - `:one`: 단일 row 반환
   - `:many`: 여러 row 반환 (슬라이스)
4. schema에 해당하는 migration 디렉터리 연결(이미 되어있음)
    ```sql
    CREATE TABLE ...
    ```
5. 코드 생성
    ```shell
    sqlc generate
    ```
    `internal/sqlc`디렉터리에 코드 생성됨
6. 기존 repo에 통합!

    repo는 쿼리를 돌려서 필요한 값을 service에 넘김.

    따라서, sqlc패키지에서 생성된 *sqlc.Queries를 필드로 쓰면 됨


## 추가적으로 고려되는것(하다보니)
### 1. Spring Boot Timeout Settings
```yaml
# Spring Boot Defaults (HikariCP)
spring.datasource.hikari.connection-timeout=30000      # 30초
spring.datasource.hikari.max-lifetime=1800000          # 30분  
spring.datasource.hikari.idle-timeout=600000           # 10분

# Transaction (관례적으로)
spring.transaction.default-timeout=30                  # 30초

# HTTP Request (Tomcat 기본은 무제한이지만, 실무에서는)
server.connection-timeout=30000                        # 30초
```