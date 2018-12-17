package main

import (
	"os"
	"testing"
)

func TestConfigReads(t *testing.T) {
	var props props
	const rackhdUrl = "http://192.168.1.165:8080"
	const configPath = "./config.test.yaml"
	err := os.Setenv(RackHdApiUrlEnvVarName, rackhdUrl)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	err = os.Setenv(AnsibleRackHdConfigPath, configPath)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	props = getPropsFromConfig()
	if props.rackhdUrl != rackhdUrl {
		t.Errorf("\n%s  \n%s", props.rackhdUrl, rackhdUrl)
	}
	if len(props.groups) != 3 {
		t.Errorf("\n%d  \n%d", len(props.groups), 3)
	}
	if props.groups[2] != "test_group_2" {
		t.Errorf("\n%s  \n%s", props.groups[2], "test_group_2")
	}
	if props.filterGroup != "something_to_filter" {
		t.Errorf("\n%s  \n%s", props.filterGroup, "something_to_filter")
	}
	err = os.Unsetenv(RackHdApiUrlEnvVarName)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	err = os.Unsetenv(AnsibleRackHdConfigPath)
	if err != nil {
		t.Errorf("%s\n", err)
	}
}

//TODO
func TestHandleList(t *testing.T) {

}
