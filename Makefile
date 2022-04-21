.PHONY: run-backend run-worker run-frontend syncdb release

MAKEFLAGS += --warn-undefined-variables

# Build variables
REGISTRY_URI :=www.haozheyu.top
RELEASE_VERSION :=$(shell git describe --always --tags)
UI_BUILD_VERSION :=v1.0.0
SERVER_BUILD_VERSION :=v1.0.0

# run module
run-backend:
	cd src/backend/ && go run main.go

# dev
syncdb:
	go run src/backend/database/syncdb.go orm syncdb

sqlall:
	go run src/backend/database/syncdb.go orm sqlall > _dev/kube_web.sql

initdata:
	go run src/backend/database/generatedata/main.go > _dev/kube_web-data.sql

swagger-openapi:
	cd src/backend && swagger generate spec -o ./swagger/openapi.swagger.json

# release, requiring Docker 17.05 or higher on the daemon and client
build-backend-image:
	@echo "version: $(RELEASE_VERSION)"
	docker build --no-cache -t $(REGISTRY_URI)/kube-web-backend:$(RELEASE_VERSION) .

push-image:
	docker tag $(REGISTRY_URI)/kube-web-backend:$(RELEASE_VERSION) $(REGISTRY_URI)/kube-web-backend:latest
	docker push $(REGISTRY_URI)/kube-web-backend:$(RELEASE_VERSION)
	docker push $(REGISTRY_URI)/kube-web-backend:latest

release: build-backend-image push-image
