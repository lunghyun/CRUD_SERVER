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

### 3. `sqlc` 도입
- [ ] ***is it Done?***


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