FROM golang:1

# Configure apt
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get -y install --no-install-recommends apt-utils dialog less 2>&1

## YOUR DOCKERFILE CONTENT GOES HERE
ENV DEBIAN_FRONTEND=dialog

# RUN go get -x -d github.com/stamblerre/gocode 2>&1 \
#     && go build -o gocode-gomod github.com/stamblerre/gocode \
#     && mv gocode-gomod $GOPATH/bin/ \
#     && go get -u \
#     golang.org/x/tools/gopls \
#     github.com/mdempsky/gocode \
#     github.com/uudashr/gopkgs/cmd/gopkgs \
#     github.com/ramya-rao-a/go-outline \
#     github.com/acroca/go-symbols \
#     golang.org/x/tools/cmd/guru \
#     golang.org/x/tools/cmd/gorename \
#     github.com/go-delve/delve/cmd/dlv \
#     github.com/stamblerre/gocode \
#     github.com/rogpeppe/godef \
#     golang.org/x/tools/cmd/goimports \
#     golang.org/x/lint/golint 2>&1
