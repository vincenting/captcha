# Install dirs auto
mkdir -p src
mkdir -p pkg
mkdir -p bin

export GOPATH="$PWD"
export PATH=$PATH:$GOPATH/bin

# install necessary dependence
echo 'Downloading necessary files'
go get github.com/gographics/imagick/imagick
echo 'Dependence install success.'

# build file
cd src/open.jianxin.io/server
echo 'Start build'
go install
cd ../../../

# move files
echo 'Moving system fils to build'
mkdir -p build
mkdir -p build/tmp
mkdir -p build/bin

cp -r src/open.jianxin.io/resource build
mv bin/server build/bin

echo 'Finish'