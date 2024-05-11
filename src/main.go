package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run" :
		parent()
	case "child":
		child()
	default:
		panic("Invalid command")
	}
}

func parent() {
	cmd := exec.Command(os.Args[0], append([]string{"child"}, os.Args[2:]...)...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR: " , err.Error())
		os.Exit(1)
	}
}

func child() {
	// must(syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""))
	// must(os.MkdirAll("rootfs/oldrootfs", 0700))
	// must(syscall.PivotRoot("rootfs", "rootfs/oldrootfs"))
	// must(os.Chdir("/"))
	fmt.Printf("child running %v as PID %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	setup()
	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR: " , err.Error())
		os.Exit(1)
	}
}

func setup() {
	if err := syscall.Sethostname([]byte("container")); err != nil {
		panic(fmt.Sprintf("Sethostname: %v\n", err))
	}

	pwd , err := os.Getwd()
	if(err != nil) {
		panic(fmt.Sprintf("Getwd Error: %v\n", err))
	
	}

	target := path.Join(pwd, "rootfs")
	if err := syscall.Chroot(target); err != nil {
		panic(fmt.Sprintf("Chroot: %v\n", err))
	}

	if err := os.Chdir("/"); err != nil {
		panic(fmt.Sprintf("Chdir: %v\n", err))
	}

	if err := syscall.Mount("proc", "proc", "proc", 0, ""); err != nil {
		log.Printf("Failed to mount /proc to %s: %v",target, err)
		panic(err)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}