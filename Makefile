all: bin/example

PLATFORM=local

.PHONY: bin/example
bin/example:
	@docker build . --target bin \
	--output bin/ \
	--platform ${PLATFORM}

# https://go.dev/doc/install/source#environment