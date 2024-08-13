PLATFORM ?= linux/amd64
TAG_PREFIX ?= registry:5000/
TAG_PREFIX_SANDBOX ?= eco-registry.supos.com/supvxcen/
DOCKER_PREFIX ?= openiiot-
DEFAULT_TAG := 1.0.0
# Define specific tags for services here, e.g., server_TAG := 1.0.1
server_TAG := 1.0.0
# Add more service-specific tags as needed

DOCKERFILE_FOLDER := deployment
DOCKERFILE_NAME := Dockerfile
SAVE_DIR ?= deployment/binary

SERVICES := server nodered grafana tdengine emqx mysql web consolemanager busybox

# Function to get a service-specific tag or default to DEFAULT_TAG
get_tag = $(or $($1_TAG),$(DEFAULT_TAG))

.PHONY: docker_all clean_all $(SERVICES) clean tag help save

docker_all: $(SERVICES)

$(SERVICES):
	$(eval SERVICE_TAG := $(call get_tag,$@,$(TAG)))
	@echo "Building with tag: $(SERVICE_TAG)"
	docker build --platform $(PLATFORM) -t $(DOCKER_PREFIX)$@:$(SERVICE_TAG) -f $(DOCKERFILE_FOLDER)/$@/$(DOCKERFILE_NAME) .

clean_all: $(addprefix clean_,$(SERVICES))

$(addprefix clean_,$(SERVICES)):
	$(eval SERVICE_TAG := $(call get_tag,$(subst clean_,,$@)))
	@echo "Stopping and cleaning service $(subst clean_,,$@)..."
	-docker ps -q --filter "ancestor=$(DOCKER_PREFIX)$(subst clean_,,$@):$(SERVICE_TAG)" | xargs -r docker stop
	-docker ps -a -q --filter "ancestor=$(DOCKER_PREFIX)$(subst clean_,,$@):$(SERVICE_TAG)" | xargs -r docker rm
	-docker rmi $(DOCKER_PREFIX)$(subst clean_,,$@):$(SERVICE_TAG)

build: docker_all
clean: clean_all

tag:
	$(foreach service,$(SERVICES),$(eval SERVICE_TAG := $(call get_tag,$(service))) docker tag $(DOCKER_PREFIX)$(service):$(SERVICE_TAG) $(TAG_PREFIX_SANDBOX)$(DOCKER_PREFIX)$(service):$(SERVICE_TAG);)

untag:
	$(foreach service,$(SERVICES),$(eval SERVICE_TAG := $(call get_tag,$(service))) docker rmi $(TAG_PREFIX_SANDBOX)$(DOCKER_PREFIX)$(service):$(SERVICE_TAG);)

push:
	$(foreach service,$(SERVICES),$(eval SERVICE_TAG := $(call get_tag,$(service))) docker push $(TAG_PREFIX_SANDBOX)$(DOCKER_PREFIX)$(service):$(SERVICE_TAG);)

save:
	@mkdir -p $(SAVE_DIR)
	$(foreach service,$(SERVICES),$(eval SERVICE_TAG := $(call get_tag,$(service))) docker save -o $(SAVE_DIR)/$(DOCKER_PREFIX)$(service).tar $(DOCKER_PREFIX)$(service):$(SERVICE_TAG);)
	
$(addprefix save_,$(SERVICES)):
	$(eval SERVICE := $(subst save_,,$@))
	$(eval SERVICE_TAG := $(call get_tag,$(SERVICE)))
	@echo "Saving service $(SERVICE) with tag $(SERVICE_TAG)..."
	@mkdir -p $(SAVE_DIR)
	docker save -o $(SAVE_DIR)/$(DOCKER_PREFIX)$(SERVICE).tar $(DOCKER_PREFIX)$(SERVICE):$(SERVICE_TAG)

help:
	@echo "1. push image to sandbox"
	@echo "docker login eco-registry.supos.com --username=zhanghongzhi@freezonex.io"
	@echo "docker push eco-registry.supos.com/supvxcen/node-red:1.0.0"
	@echo "2. push image to http://47.236.10.165/"
	
