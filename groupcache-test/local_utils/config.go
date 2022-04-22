package local_utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	GCFG *SrvCFG
)

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
		Id   string `yaml:"id"`
		Addr string `yaml:"addr"`
	}
}
