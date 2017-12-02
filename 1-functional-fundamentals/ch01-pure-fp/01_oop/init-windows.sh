#!/bin/bash -e

echo "Installing goimports..."
export GOPATH=C:\\Go\\src\\
rm -rf $GOPATH\\golang.org\\x\\tools 2>/dev/null
mkdir -p $GOPATH\\golang.org\\x\\tools
cd $GOPATH\\golang.org\\x\\tools
git clone https://go.googlesource.com/tools .
cd $GOPATH
cd golang.org/x/tools/cmd/goimports
go build
cp ./goimports.exe /usr/local/bin

echo "Installing glide..."
cd $GOPATH
rm -rf github.com/Masterminds/glide 2>/dev/null
mkdir -p github.com/Masterminds/glide
cd github.com/Masterminds/glide
git clone https://github.com/Masterminds/glide.git .
go build
cp ./glide.exe /usr/local/bin/

echo "Installed goimports and $(glide --version)"
