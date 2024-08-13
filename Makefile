ifeq ($(OS),Windows_NT)
CUR_DIR=$(shell echo %CD%)
else
CUR_DIR=$(shell pwd)
endif

IMAGE=tax_api
TAG=latest
RELEASE_NAME=tax-api
DC_FILE=-f ${CUR_DIR}/deployment/docker-compose.yaml

compile:
	docker build --no-cache -f .docker/Dockerfile -t ${IMAGE}:${TAG} --target builder .

deploy:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} up -d

deploy-app:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} up -d app

deploy-postgres:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} up -d postgres_db postgres_migrate

deploy-redis:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} up -d redis_db

delete:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} rm -sf

delete-app:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} rm -sf app

delete-postgres:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} rm -sf postgres_db postgres_migrate

delete-redis:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} rm -sf redis_db