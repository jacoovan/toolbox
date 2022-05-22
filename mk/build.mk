.PHONY: build docker push

build:
	$(info ==================================================)
	$(info stage: build$(NEWLINE))
	$(info ==================================================)
	# 1.
	go mod tidy
	# 2.
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN_FILE) $(ENTRY_GO_FILE)

docker:
	$(info ==================================================)
	$(info stage: docker$(NEWLINE))
	$(info NEW_IMAGE: ${IMAGE_NAME}:$(IMAGE_TAG))
	$(info ==================================================)
	# 1.
	chmod a+x $(BIN_FILE)
	# 2.
	docker build -f $(DOCKER_FILE) -t ${IMAGE_NAME}:$(GIT_COMMIT) .
	# 3.
	docker tag ${IMAGE_NAME}:$(GIT_COMMIT) ${IMAGE_NAME}:$(IMAGE_TAG)

push:
	$(info ==================================================)
	$(info stage: push$(NEWLINE))
	$(info NEW_IMAGE: ${IMAGE_NAME}:$(IMAGE_TAG))
	$(info ==================================================)
	# 1.
	docker push ${IMAGE_NAME}:$(GIT_COMMIT)
	# 2.
	docker push ${IMAGE_NAME}:$(IMAGE_TAG)

