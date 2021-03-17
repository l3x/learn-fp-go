FROM golang:1.15
MAINTAINER Michele Bertasi
MAINTAINER Aleksey Kislitsa

ADD fs/ /

# install pagkages
RUN apt-get update                                                      && \
    apt-get install -y ncurses-dev libtolua-dev exuberant-ctags gdb     && \
    ln -s /usr/include/lua5.2/ /usr/include/lua                         && \
    ln -s /usr/lib/x86_64-linux-gnu/liblua5.2.so /usr/lib/liblua.so     && \
# cleanup
    apt-get clean && rm -rf /var/lib/apt/lists/*

# build and install vim
RUN cd /tmp                                                             && \
    git clone --depth 1 https://github.com/vim/vim.git                  && \
    cd vim                                                              && \
    ./configure --with-features=huge --enable-luainterp                    \
        --enable-gui=no --without-x --prefix=/usr                       && \
    make VIMRUNTIMEDIR=/usr/share/vim/vim82                             && \
    make install                                                        && \
# cleanup
    rm -rf /tmp/* /var/tmp/*

# get go tools
RUN go get golang.org/x/tools/cmd/godoc                                 && \
    go get github.com/nsf/gocode                                        && \
    go get golang.org/x/tools/cmd/goimports                             && \
    go get github.com/rogpeppe/godef                                    && \
    go get golang.org/x/tools/cmd/gorename                              && \
    go get golang.org/x/lint/golint                                     && \
    go get github.com/kisielk/errcheck                                  && \
    go get github.com/jstemmer/gotags                                   && \
    go get github.com/tools/godep                                       && \
    GO111MODULE=on go get golang.org/x/tools/gopls@latest               && \
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
