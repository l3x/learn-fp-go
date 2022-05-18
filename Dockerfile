FROM golang:1.18
LABEL MAINTAINER Michele Bertasi
LABEL MAINTAINER Aleksey Kislitsa

ADD fs/ /
ADD go-env /usr/local/bin
ADD go-init /usr/local/bin
ADD go-ide /usr/local/bin

# install pagkages
RUN apt-get update                                                      && \
    apt-get install -y ncurses-dev libtolua-dev                            \
    exuberant-ctags gdb vim-nox                                         && \
    ln -s /usr/include/lua5.2/ /usr/include/lua                         && \
    ln -s /usr/lib/x86_64-linux-gnu/liblua5.2.so /usr/lib/liblua.so     && \
# cleanup
    apt-get clean && rm -rf /var/lib/apt/lists/*

# get go tools
RUN go install golang.org/x/tools/cmd/godoc@latest                      && \
    go install github.com/nsf/gocode@latest                             && \
    go install golang.org/x/tools/cmd/goimports@latest                  && \
    go install github.com/rogpeppe/godef@latest                         && \
    go install golang.org/x/tools/cmd/gorename@latest                   && \
    go install golang.org/x/lint/golint@latest                          && \
    go install github.com/kisielk/errcheck@latest                       && \
    go install github.com/jstemmer/gotags@latest                        && \
    go install github.com/tools/godep@latest                            && \
    go install golang.org/x/tools/gopls@latest                          && \
    mv /go/bin/* /usr/local/go/bin                                      && \
# cleanup
    rm -rf /go/src/* /go/pkg

# add dev user
RUN adduser dev --disabled-password --gecos ""                          && \
    echo "ALL            ALL = (ALL) NOPASSWD: ALL" >> /etc/sudoers     && \
    chown -R dev:dev /home/dev /go

USER dev
ENV HOME /home/dev
WORKDIR /home/dev

# install vim plugins
RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim && \
    vim +PlugInstall +qall
