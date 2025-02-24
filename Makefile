ifeq ($(OS),Windows_NT)
CUR_DIR=$(shell echo %CD%)
else
CUR_DIR=$(shell pwd)
endif

IMAGE=api
TAG=latest
RELEASE_NAME=api
DC_FILE=-f ${CUR_DIR}/deployment/docker-compose.yaml

.PHONY: compile copy-env copy-env-windows deploy deploy-app deploy-postgres delete delete-app delete-postgres

compile:
	docker build --no-cache -f .docker/Dockerfile -t ${IMAGE}:${TAG} --target builder .

copy-env:
	cp deployment/.env.example deployment/.env

copy-env-windows:
	copy deployment\.env.example deployment\.env

deploy:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} up -d

deploy-app:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} up -d app

deploy-postgres:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} up -d postgres_db postgres_migrate

delete:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} rm -sf

delete-app:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} rm -sf app

delete-postgres:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} rm -sf postgres_db postgres_migrate
