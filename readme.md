# Template
https://www.docker.com/blog/containerize-your-go-developer-environment-part-1/

bash: export DOCKER_BUILDKIT=1
pwsh: $env:DOCKER_BUILDKIT=1
cmd: set DOCKER_BUILDKIT=1

make
make PLATFORM=windows/amd64
make PLATFORM=linux/arm64

# Implementation
https://blog.logrocket.com/configuring-the-go-http-client/