GCC=go
GCMD=run
GPATH=main.go
GO_VERSION=1.10

run:
	make build
	$(GCC) $(GCMD) $(GPATH)

build:
	make build_db

build_db:
	rm pkg/db/db_structs.go
	go run pkg/db/main.go -json=./pkg/db/config.json
	mv db_structs.go pkg/db/

install:
	make install_routes
	make install_db

install_routes:
	go get -u github.com/gorilla/mux
install_db:
	go get -u github.com/go-xorm/xorm

# Please add these to your bash file
#
# export GOROOT=/usr/local/go
# export PATH=$PATH:/usr/local/go/bin
# export GOPATH=/home/$USER/Projects/Go
# export PATH=$PATH:$GOPATH/bin
install_go:
	curl https://dl.google.com/go/go$(GO_VERSION).linux-amd64.tar.gz > go$(GO_VERSION).linux-amd64.tar.gz
	tar -xvf go$(GO_VERSION).linux-amd64.tar.gz
	sudo rm -rf /usr/local/go
	sudo mv go /usr/local
	rm go$(GO_VERSION).linux-amd64.tar.gz