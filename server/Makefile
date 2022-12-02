run:
	go run cmd/server/main.go

db:
ifeq ($(c),)
	echo "Choose up, or down"
else ifeq ($(c), up)
	docker-compose -f .docker/docker-compose.dev.yaml $(c) --build -d
else ifeq ($(c), down)
	docker-compose -f .docker/docker-compose.dev.yaml $(c) -v
else ifeq ($(c), init)
	mysql -h127.0.0.1 -uroot -pkiara_is_not_a_vtuber < tools/init/db.sql
endif

create_migration:
ifneq ($(name),)
	cd tools/migrations && goose create $(name) sql
else
	echo "name is empty"
endif

goose:
ifeq ($(env), $(filter $(env),local dev))
	goose -dir tools/migrations mysql root:kiara_is_not_a_vtuber@tcp\(localhost:3306\)/kiaranote?parseTime=true $(c)
else ifeq ($(env), test)
	goose -dir tools/migrations mysql root:kiara_is_not_a_vtuber@tcp\(localhost:3306\)/kiaranote_test?parseTime=true $(c)
endif
