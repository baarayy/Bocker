config:
		mkdir rootfs
		tar -xvf rootfs.tar -C rootfs
build:
		go build main.go
run: 
		sudo ./container run /bin/bash