package config

import (
	"io/ioutil"
	"testing"
)

func TestConfig(t *testing.T) {
	configFile := "config.ini"
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		t.Error("read file failed")
		return
	}
	var result *ConfigInfo = &ConfigInfo{}
	//var testint *int
	err = UnMarshal(data, result)
	if err != nil {
		t.Error("unmashal filed Error:", err)
	}
	t.Logf("unmashal success,conf: %#v", result)
}
