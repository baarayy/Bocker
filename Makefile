config:
	mkdir -p rootfs
	sudo tar -xvf rootfs.tar -C rootfs

build: 
	GOOS=linux GOARCH=amd64 go build

run: 
	sudo ./bocker run /bin/bash 

.PHONY: all
all: config build run