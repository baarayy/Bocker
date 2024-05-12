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
	case "run":
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
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		panic("start parent error" + err.Error())
	}
	log.Printf("container PID: %d", cmd.Process.Pid)
	if err := putIface(cmd.Process.Pid); err != nil {
		panic("putIface error" + err.Error())
	}
	if err := cmd.Wait(); err != nil {
		panic("wait error" + err.Error())
	}
}

func child() {
	fmt.Printf("child running %v as PID %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	setup()
	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
}

func setup() {
	if err := syscall.Sethostname([]byte("bocker")); err != nil {
		panic(fmt.Sprintf("Sethostname: %v\n", err))
	}

	pwd, err := os.Getwd()
	if err != nil {
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
		log.Printf("Failed to mount /proc to %s: %v", target, err)
		panic(err)
	}
	lnk, err := waitForIfac()
	if err != nil {
		panic(fmt.Sprintf("waitForIfac error: %v\n", err))
	}
	if err := setupIface(lnk); err != nil {

	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
