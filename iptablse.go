package main

import (
	"fmt"

	"github.com/coreos/go-iptables/iptables"
)

type IPTablesRule struct {
	table    string
	chain    string
	rulespec []string
}

func getIptableRules(n, hostDevice, virtualDevice string) []IPTablesRule {
	return []IPTablesRule{
		{"nat", "POSTROUTING", []string{"-s", n, "-o", hostDevice, "-j", "MASQUERADE"}},
		{"filter", "FORWARD", []string{"-i", hostDevice, "-o", virtualDevice, "-j", "ACCEPT"}},
		{"filter", "FORWARD", []string{"-i", virtualDevice, "-o", hostDevice, "-j", "ACCEPT"}},
	}
}

func setIptables(rules []IPTablesRule) error {
	ipt, err := iptables.New()
	if err != nil {
		return fmt.Errorf("Error initializing iptables: %v", err)
	}
	for _, rule := range rules {
		if err := ipt.AppendUnique(rule.table, rule.chain, rule.rulespec...); err != nil {
			return fmt.Errorf("Error appending iptables rule: %v", err)
		}
	}
	return nil
}

func teardownIPTables(ipt iptables.IPTables, rules []IPTablesRule) {
	for _, rule := range rules {
		ipt.Delete(rule.table, rule.chain, rule.rulespec...)
	}
}
