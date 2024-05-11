package main

import (
	"fmt"
	"os"
	"os/exec"
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

	if err := syscall.Sethostname([]byte("container")); err != nil {
		panic(fmt.Sprintf("Sethostname: %v", err))
	}

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR: " , err.Error())
		os.Exit(1)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}