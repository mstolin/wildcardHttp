PORT = 5000
NAME = wildcardhttp

ifeq ($(shell command -v podman 2> /dev/null),)
    DOCKER=docker
else
    DOCKER=podman
endif

.PHONY: build
build:
	go build -o whttp main.go

.PHONY: run
run:
	./whttp :$(PORT)

.PHONY: build-docker
build-docker:
	$(DOCKER) build -t $(NAME) .

.PHONY: run-docker
run-docker:
	$(DOCKER) run --rm -d -p 127.0.0.1:$(PORT):80 --name $(NAME) localhost/$(NAME):latest

.PHONY: stop-docker
stop-docker:
	$(DOCKER) stop $(NAME)
