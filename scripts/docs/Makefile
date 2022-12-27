# Documentation

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

TARGET_MAX_CHAR_NUM=40

## show help
help: Makefile
	@awk '/^[a-zA-Z\-_0-9\/ ]+:/ { \
  		helpMessage = match(lastLine, /^## (.*)/); \
  		if (helpMessage) { \
  			helpCommand = $$1; sub(/:$$/, "", helpCommand); \
  			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
  			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
  		} \
  	} \
  	{ lastLine = $$0 }' $(MAKEFILE_LIST)
