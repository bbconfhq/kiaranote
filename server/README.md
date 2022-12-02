# kiaranote-server 

## 로컬 개발 환경 설정

### Prerequisites

- Docker and docker-compose
- Golang 1.18+
- [goose](https://github.com/pressly/goose)

### 프로젝트 의존성 설치

```shell
# 프로젝트 의존성 실행
go mod tidy

# swag 설치
go install github.com/swaggo/swag/cmd/swag@latest
```

### Database

**`x86_64` 아키텍쳐 환경이 아닌 경우에는 따로 mysql 이미지를 내려받아야합니다.**

```shell
docker pull mysql:8.0 --platform=x86_64
```

`make db c=[up|down|init]` 명령어로 사용합니다.

**Example:**

새 MySQL Docker 생성

```shell
make db c=up
```

MySQL Docker 제거 및 Volume 제거

```shell
make db c=down
```

초기 DB 스크립트 실행 (Non-docker 환경에서 필요)

```shell
make db c=init
```

#### Database Migration

`make goose env=[local|dev|test] c=[goose command]` 명령어로 사용합니다.

**Example:**

새 마이그레이션 생성

```shell
make create_migration name=<migration name>
```

마이그레이션 적용

```shell
make goose env=local c=up
```

마지막 마이그레이션 롤백

```shell
make goose env=local c=down
```

### Server

서버를 실행합니다.

```shell
make run
```

### Swagger

현재 swagger를 사용하려면 아래 명령어 입력 후 서버를 실행해주어야 합니다.

`docs`를 커밋에 올리지 말아주세요.

> Note: `--parseDependency` 옵션이 에러가 발생합니다. 원인 파악 후 조치하겠습니다.

```shell
swag init -d ./cmd/server/,./ --parseInternal --generatedTime
```
