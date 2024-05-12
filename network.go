package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/vishvananda/netlink"
)
const (
	bridgeName = "container0"
	vethPrefix = "veth-pair"
	ipAddr     = "10.88.37.1/24"
)
func createBridge() error {
	_,err := net.InterfaceByName(bridgeName)
	if err == nil {
		return nil
	}
	la := netlink.NewLinkAttrs()
	la.Name = bridgeName
	br := &netlink.Bridge{LinkAttrs: la}

	if err := netlink.LinkAdd(br); err != nil {
		return fmt.Errorf("Bridge create failed: %v", err)
	}

	addr , err := netlink.ParseAddr(ipAddr)
	if err != nil {
		return fmt.Errorf("ParseAddr failed: %v", err)
	}

	if err := netlink.AddrAdd(br, addr); err != nil {
		return fmt.Errorf("AddrAdd failed: %v", err)
	}

	if err := netlink.LinkSetUp(br); err != nil {
		return err
	}
	return nil
}
func createVethPair(pid int) error {
	br , err := netlink.LinkByName(bridgeName)
	if err != nil {
		return err
	}
	x1,x2 := rand.Intn(1000),rand.Intn(1000)
	parentName := fmt.Sprintf("%s%d" , vethPrefix , x1)
	peerName := fmt.Sprintf("%s%d" , vethPrefix , x2)

	la := netlink.NewLinkAttrs()
	la.Name = parentName
	la.MasterIndex = br.Attrs().Index
	
	vp := &netlink.Veth {
		LinkAttrs: la,
		PeerName: peerName,
	}
	if err := netlink.LinkAdd(vp); err != nil {
		return fmt.Errorf("LinkAdd failed: %v", err)
	}
	peer,err := netlink.LinkByName(peerName)
	if err != nil {
		fmt.Errorf("LinkByName failed: %v", err)
	}
	if err := netlink.LinkSetNsPid(peer,pid); err != nil {
		return fmt.Errorf("LinkSetNsPid failed: %v", err)
	}
	if err := netlink.LinkSetUp(vp); err != nil {
		return fmt.Errorf("LinkSetUp failed: %v", err)
	
	}
	return nil
}

func putIface(pid int) error {
	if err := createBridge(); err != nil {
		return fmt.Errorf("createBridge failed: %v", err)
	}
	if err := createVethPair(pid); err != nil {
		return fmt.Errorf("createVethPair failed: %v", err)
	}
	return nil
}

func waitForIfac() (netlink.Link , error) {
	log.Printf("Waiting for %s network interface to appear...", bridgeName)
	start := time.Now()
	for {
		fmt.Print(".")
		if time.Since(start) > 5 * time.Second {
			fmt.Printf("\n")
			return nil , fmt.Errorf("failed to find venth interface in 5 seconds")
		}
		lst , err := netlink.LinkList()
		if err != nil {
			fmt.Printf("\n")
			return nil , err
		}
		for _,l := range lst {
			if l.Type() == "veth" {
				fmt.Printf("\n")
				return l , nil
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}