# kiaranote-server 

* [Concept](https://demo.hedgedoc.org/JOieJRo0QNesVBdgEel48Q)
* [ERD](https://www.erdcloud.com/d/h2HrcZe3gYPNMe6Zb)

## 로컬 개발 환경 설정

### Prerequisites

- Docker and docker-compose
- Golang 1.18+
- [goose](https://github.com/pressly/goose)

### Project Setup

프로젝트를 실행시키기 위해서는 아래 명령어를 따라야 합니다.

```shell
# Install swag
go install github.com/swaggo/swag/cmd/swag@latest

# Install goose
go install github.com/pressly/goose/v3/cmd/goose@latest

# Generate swagger docs
swag init -d ./cmd/server/,./ --parseInternal --generatedTime

# Install dependencies
go mod tidy

# Use mysql x86_64 (also support arm64)
docker pull mysql:8.0 --platform=x86_64

# Run docker-compose
make db c=up

# Run migration
make goose env=local c=up

# Create .env file (can copy from .env.example)
cp .env.example .env

# Run server
make run
```

### Makefile

Makefile에는 자주 사용하는 명령어가 정의되어 있습니다.

#### Database

Docker를 사용하여 MySQL을 실행합니다.

- MySQL Docker 실행

```shell
make db c=up
```

- MySQL Docker 제거 및 Volume 제거

```shell
make db c=down
```

- 초기 DB 스크립트 실행 (Non-docker 환경)

```shell
make db c=init
```

#### Migration

`goose`를 사용하여 마이그레이션을 실행합니다.

- 마이그레이션 생성

```shell
make create_migration name=<migration name>
```

- 마이그레이션 실행

```shell
make goose env=local c=up
```

- 마이그레이션 롤백

```shell
make goose env=local c=down
```

#### Server

`make run`

- 서버 실행

```shell
make run
```

### Swagger

> Note: `--parseDependency` 옵션이 에러가 발생합니다. 원인 파악 후 조치하겠습니다.

swagger 문서는 아래 명령어 입력 시 자동으로 생성됩니다. 

서버 실행 후 `/swagger/index.html` 에 접속하여 확인 가능합니다.

```shell
swag init -d ./cmd/server/,./ --parseInternal --generatedTime
```

### Test

테스트를 위해선 데이터베이스 연결(`.env.test` 참고)이 필요합니다.

```shell
go test ./...
```