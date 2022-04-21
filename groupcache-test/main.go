package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	GCFG *SrvCFG
)

func main() {

}

func Logf(vals ...interface{}) {
	outs := make([]interface{}, 0, len(vals)+2)
	outs = append(outs, fmt.Sprintf("[%s] [%s] ", time.Now().String(), GCFG.Server.Id))
	outs = append(outs, vals...)
	fmt.Println(outs...)
}

func LoadCFG(file string) (*SrvCFG, error) {
	fs, fsErr := ioutil.ReadFile(file)
	if fsErr != nil {
		return nil, fsErr
	}
	var retCfg SrvCFG
	parseErr := yaml.Unmarshal(fs, &retCfg)
	if parseErr != nil {
		return nil, parseErr
	}
	GCFG = &retCfg
	return &retCfg, nil
}

type SrvCFG struct {
	Server struct {
		Id string `yaml:"id"`
	}
}
