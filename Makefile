PLATFORM ?= linux/amd64
TAG_PREFIX ?= registry:5000/
TAG_PREFIX_SANDBOX ?= eco-registry.supos.com/supvxcen/
DOCKER_PREFIX ?= openiiot_
TAG := 1.0.0
DOCKERFILE_FOLDER := deployment
DOCKERFILE_NAME := Dockerfile
SAVE_DIR ?= deployment/binary

SERVICES := server nodered grafana tdengine emqx mysql web consolemanager
.PHONY: docker_all clean_all $(SERVICES) clean tag help save

docker_all: $(SERVICES)

$(SERVICES):
	docker build --platform $(PLATFORM) -t $(DOCKER_PREFIX)$@:$(TAG) -f $(DOCKERFILE_FOLDER)/$@/$(DOCKERFILE_NAME) .

clean_all: $(addprefix clean_,$(SERVICES))

$(addprefix clean_,$(SERVICES)):
	@echo "Stopping and cleaning service $(subst clean_,,$@)..."
	-docker ps -q --filter "ancestor=$(DOCKER_PREFIX)$(subst clean_,,$@):$(TAG)" | xargs -r docker stop
	-docker ps -a -q --filter "ancestor=$(DOCKER_PREFIX)$(subst clean_,,$@):$(TAG)" | xargs -r docker rm
	-docker rmi $(DOCKER_PREFIX)$(subst clean_,,$@):$(TAG)

build: docker_all
clean: clean_all

# tag:
# 	$(foreach service,$(SERVICES),docker tag $(DOCKER_PREFIX)$(service):$(TAG) $(TAG_PREFIX)$(DOCKER_PREFIX)$(service):$(TAG);)

# untag:
# 	$(foreach service,$(SERVICES),docker rmi $(TAG_PREFIX)$(DOCKER_PREFIX)$(service):$(TAG);)

tag:
	$(foreach service,$(SERVICES),docker tag $(DOCKER_PREFIX)$(service):$(TAG) $(TAG_PREFIX_SANDBOX)$(DOCKER_PREFIX)$(service):$(TAG);)

untag:
	$(foreach service,$(SERVICES),docker rmi $(TAG_PREFIX_SANDBOX)$(DOCKER_PREFIX)$(service):$(TAG);)

# only push to sandbox, private installation not use this command
push:
	$(foreach service,$(SERVICES),docker push $(TAG_PREFIX_SANDBOX)$(DOCKER_PREFIX)$(service):$(TAG);)
	
help:
	@echo "1. push image to sandbox"
	@echo "docker login eco-registry.supos.com --username=zhanghongzhi@freezonex.io"
	@echo "docker push eco-registry.supos.com/supvxcen/node-red:1.0.0"
	@echo "2. push image to http://47.236.10.165/"

# New target to save images
save:
	@mkdir -p $(SAVE_DIR)
	$(foreach service,$(SERVICES),docker save -o $(SAVE_DIR)/$(DOCKER_PREFIX)$(service).tar $(DOCKER_PREFIX)$(service):$(TAG);)

$(addprefix save_,$(SERVICES)):
	@echo "Saving service $(subst save_,,$@)..."
	@mkdir -p $(SAVE_DIR)
	docker save -o $(SAVE_DIR)/$(DOCKER_PREFIX)$(subst save_,,$@).tar $(DOCKER_PREFIX)$(subst save_,,$@):$(TAG)
