ARG base
FROM ${base}

USER gitpod
ENV HOME=/home/gitpod
WORKDIR $HOME
ENV GO_VERSION=1.14 \
    GOPATH=$HOME/go-packages \
    GOROOT=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH
RUN curl -fsSL https://storage.googleapis.com/golang/go$GO_VERSION.linux-amd64.tar.gz | tar xzs
# install VS Code Go tools from https://github.com/Microsoft/vscode-go/blob/0faec7e5a8a69d71093f08e035d33beb3ded8626/src/goInstallTools.ts#L19-L45
RUN go get -u -v \
        github.com/mdempsky/gocode \
        github.com/uudashr/gopkgs/cmd/gopkgs \
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
        github.com/alecthomas/gometalinter \
        honnef.co/go/tools/... \
        # github.com/golangci/golangci-lint/cmd/golangci-lint \
        github.com/mgechev/revive \
        github.com/sourcegraph/go-langserver \
        github.com/go-delve/delve/cmd/dlv \
        github.com/davidrjenni/reftools/cmd/fillstruct \
        github.com/godoctor/godoctor && \
    GO111MODULE=on go get -u -v \
        golang.org/x/tools/gopls@latest && \
    go get -u -v -d github.com/stamblerre/gocode && \
    go build -o $GOPATH/bin/gocode-gomod github.com/stamblerre/gocode && \
    rm -rf $GOPATH/src && \
    sudo rm -rf $GOPATH/pkg && \
    rm -rf /home/gitpod/.cache/go
# user Go packages
ENV GOPATH=/workspace/go \
    PATH=/workspace/go/bin:$PATH
