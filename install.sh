#!/bin/sh

# Install dirs auto
mkdir -p src
mkdir -p pkg
mkdir -p bin

OLD_GOPATH=$GOPATH
export GOPATH=$(cd `dirname $0`; pwd)

# install necessary dependence
echo 'Downloading necessary files'
go get github.com/gographics/imagick/imagick
echo 'Dependence install success.'

# build file
cd src/captcha/server/
echo 'Start build'
go install
cd ../../../
go clean

# move files
echo 'Moving system fils to build'
mkdir -p build
mkdir -p build/tmp
mkdir -p build/bin

cp src/captcha/config.json build
cp -r src/captcha/assets build
mv bin/server build/bin/captcha

export GOPATH=$OLD_GOPATH
echo 'Finish'