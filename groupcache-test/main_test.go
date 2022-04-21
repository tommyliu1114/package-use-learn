package main

import (
	"fmt"
	"testing"
)

func TestLoadCFG(t *testing.T) {
	cfg, cfgErr := LoadCFG("./config-template.yml")
	if cfgErr != nil {
		fmt.Println(cfgErr.Error())
		return
	}
	fmt.Printf("cfg: (%+v) \n", cfg)
}

func TestLog(t *testing.T) {
	_, err := LoadCFG("./config-template.yml")
	if err != nil {
		fmt.Printf("parse err %s \n", err.Error())
		return
	}
	Logf(20, "hello world")
}
