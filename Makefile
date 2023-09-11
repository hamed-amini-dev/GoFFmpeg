#
# ─── GLOBAL VARIABLES ───────────────────────────────────────────────────────────
#
VERSION ?= $(shell echo $(shell git describe --tags) | sed 's/^v//')
OUTPUT_PATH ?= ./bin


#
# ────────────────────────────────────────────────  ──────────
#   :::::: H E L P : :  :   :    :     :        :          :
# ──────────────────────────────────────────────────────────
#
.PHONY: help

help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

#
# ──────────────────────────────────────────────────────── I ──────────
#   :::::: C O M M A N D S : :  :   :    :     :        :          :
# ──────────────────────────────────────────────────────────────────
#

FOO = $(shell docker ps -aqf name=local-irisconverter) 
FFF = $(shell docker images -q iriscontainer/irisconverter)
.build: ## build go project
    
## create go app with golang image container	
	@docker run --rm -e GOOS=$(GOOS) -e GOARCH=$(GOARCH) -v "${CURDIR}:/usr/src/app" -w "/usr/src/app" golang:latest /bin/bash -c "go build -o $(OUTPUT_PATH) ./main.go"
## create container from lasted image irisconverter
	@$(MAKE) .build-run-container
	
## copy this compiled file on container from last  image irisconverter
	@$(MAKE) .build-cp
	
## Commit this container in new image and push to docker hub 
	@$(MAKE) .build-push
	
 
## delete olded container && image irisconverter from local	
	@$(MAKE) .build-delete1
	@$(MAKE) .build-delete2
	

.build-run-container:
	@docker run --name local-irisconverter -t -d iriscontainer/irisconverter:latest 
.build-cp:
	@docker cp "${CURDIR}/bin/main" "local-irisconverter:/home/converter"
.build-push:
	@docker commit -m "commit irisconverter" -a "hamedamini" $(FOO) local-irisconverter-new-image
	@docker tag local-irisconverter-new-image:latest iriscontainer/irisconverter:latest
	@docker push iriscontainer/irisconverter:latest
.build-delete1:	
	@docker stop $(FOO)
	@docker container rm $(FOO)
	@docker image rm local-irisconverter-new-image
.build-delete2:	
	@docker image rm $(FFF)


.run: ## run convert file from input folder and result on output folder
	@docker run --rm -v "${CURDIR}/bin/EncodeFile:/home/EncodeFile" -w "/home" iriscontainer/irisconverter:latest /bin/bash -c "./converter --from * --to mp4 --t 0.4,0.6"





