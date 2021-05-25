ARG base
FROM ${base}

USER gitpod
ENV GO_VERSION=1.16.2
ENV GOPATH=$HOME/go-packages
ENV GOROOT=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH
RUN curl -fsSL https://storage.googleapis.com/golang/go$GO_VERSION.linux-amd64.tar.gz | tar xzs && \
# install VS Code Go tools from https://github.com/Microsoft/vscode-go/blob/0faec7e5a8a69d71093f08e035d33beb3ded8626/src/goInstallTools.ts#L19-L45
    go get -v \
        github.com/mdempsky/gocode \
        github.com/uudashr/gopkgs/cmd/gopkgs@v2 \
        github.com/ramya-rao-a/go-outline \
        github.com/acroca/go-symbols \
        golang.org/x/tools/cmd/guru \
        golang.org/x/tools/cmd/gorename \
        github.com/fatih/gomodifytags \
        github.com/haya14busa/goplay/cmd/goplay \
        github.com/josharian/impl \
        github.com/tylerb/gotype-live \
        github.com/rogpeppe/godef \
        github.com/zmb3/gogetdoc \
        golang.org/x/tools/cmd/goimports \
        github.com/sqs/goreturns \
        winterdrache.de/goformat/goformat \
        golang.org/x/lint/golint \
        github.com/cweill/gotests/... \
        honnef.co/go/tools/... \
        github.com/mgechev/revive \
        github.com/sourcegraph/go-langserver \
        github.com/go-delve/delve/cmd/dlv \
        github.com/davidrjenni/reftools/cmd/fillstruct \
        github.com/godoctor/godoctor && \
    # NOTE: Do not use the -u flag, as it will update your dependencies to incompatible versions.
    # see: https://github.com/golang/tools/blob/master/gopls/README.md#installation
    GO111MODULE=on go get -v \
        golang.org/x/tools/gopls@v0.6.11 && \
    sudo rm -rf $GOPATH/src $GOPATH/pkg /home/gitpod/.cache/go /home/gitpod/.cache/go-build
# user Go packages
ENV GOPATH=/workspace/go \
    PATH=/workspace/go/bin:$PATH
