# kiaranote server 

## 로컬 개발 환경 설정

### Prerequisites

- Docker and docker-compose
- Golang 1.18+
- [goose](https://github.com/pressly/goose)

### 프로젝트 의존성 설치

`go mod tidy`

### Database

**`x86_64` 아키텍쳐 환경이 아닌 경우에는 따로 mysql 이미지를 내려받아야합니다.**

```
docker pull mysql:8.0 --platform=x86_64
```

1. `make db`로 로컬 데이터베이스 컨테이너를 실행합니다.
2. `make dbinit`명령어로 기본 데이터베이스 스키마를 생성합니다. (로컬 데이터베이스 root 계정 비밀번호는 `.docker/docker-compose.dev.yaml` 참조)

#### Database Migration

`make goose env=[ENV] c=[goose command]` 명령어로 사용합니다.

**Example:**

새 마이그레이션 생성
- `make create_migration name=<migration name>`

마이그레이션 적용
- `make goose env=local c=up`

마지막 마이그레이션 롤백
- `make goose env=local c=down`

### Server

`make run`으로 서버를 실행합니다.
