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

delete:
	cd deployment && docker-compose ${DC_FILE} -p ${RELEASE_NAME} rm -sf