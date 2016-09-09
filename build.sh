export GOPATH=`pwd`
git submodule update --init --recursive
go clean
go run src/main/main.go

