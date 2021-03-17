FROM golang:1.9.2

#RUN apt update -y && apt install sudo -y
RUN apt update -y && apt install sudo git vim -y
RUN wget https://github.com/Masterminds/glide/releases/download/v0.13.3/glide-v0.13.3-linux-amd64.tar.gz -O /tmp/glide.tar.gz && \
tar -C /tmp -xzf /tmp/glide.tar.gz && cp /tmp/linux-amd64/glide /usr/local/bin

# Create non-root user
RUN useradd -m dev && \
  adduser dev sudo && \
  echo "dev:dev" | chpasswd

# Initilise base user
USER dev
WORKDIR /home/dev
ENV HOME /home/dev
ENV PATH=/home/dev/go/bin:$PATH
ENV GOPATH=/home/dev/go
ENV GOROOT=/usr/local/go
RUN go get golang.org/x/tools/cmd/goimports && \
git clone https://github.com/fatih/vim-go.git ~/.vim/pack/plugins/start/vim-go && \
git clone https://github.com/scrooloose/nerdtree.git ~/.vim/pack/dist/start/nerdtree && \
git clone https://github.com/vim-airline/vim-airline ~/.vim/pack/dist/start/vim-airline && \
git clone https://github.com/tpope/vim-fugitive.git ~/.vim/pack/dist/start/vim-fugitive
COPY .vimrc /home/dev/.vimrc
