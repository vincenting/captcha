#!/bin/sh

OLD_GOPATH=$GOPATH
export GOPATH=$(cd `dirname $0`; pwd)

# build file
cd src/captcha/server/
echo 'Start build'
go install
cd ../../../
go clean

mv bin/server build/bin/captcha

export GOPATH=$OLD_GOPATH
echo 'Finish'