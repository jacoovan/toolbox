.PHONY: deploy-init pull up down

deploy-init:
	$(info ==================================================)
	$(info stage: deploy-init$(NEWLINE))
	$(info ==================================================)
	rm -f $(ENV_FILE)
	echo COMPOSE_PROJECT_NAME=$(COMPOSE_PROJECT_NAME) >> $(ENV_FILE)
	echo IMAGE_NAME=$(IMAGE_NAME) >> $(ENV_FILE)
	echo IMAGE_TAG=$(IMAGE_TAG) >> $(ENV_FILE)
	echo SERVICE_NAME=$(SERVICE_NAME) >> $(ENV_FILE)
	echo HOST=$(HOST) >> $(ENV_FILE)
	echo PORT_WEB=$(PORT_WEB) >> $(ENV_FILE)
	echo PORT_WS=$(PORT_WS) >> $(ENV_FILE)

pull:
	$(info ==================================================)
	$(info stage: pull$(NEWLINE))
	$(info ==================================================)
	# 1.
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_FILE) pull

up:
	$(info ==================================================)
	$(info stage: up$(NEWLINE))
	$(info IMAGE: ${IMAGE_NAME}:$(IMAGE_TAG))
	$(info ==================================================)
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_FILE) up -d

down:
	$(info ==================================================)
	$(info stage: down$(NEWLINE))
	$(info IMAGE: ${IMAGE_NAME}:$(IMAGE_TAG))
	$(info ==================================================)
	docker-compose --env-file $(ENV_FILE) -f $(COMPOSE_FILE) down -v