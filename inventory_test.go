package main

import (
	"os"
	"testing"
)

func TestConfigReads(t *testing.T) {
	var props props
	const rackhdUrl 	=  "http://192.168.1.165:8080"
	const configPath	=  "./config.test.yaml"
	// Setting RACK_HD_API_URL and ANSIBLE_RACKHD_CONFIG_PATH
	// as environment variables and checking if getPropsFromConfig is reading them
	os.Setenv(RackHdApiUrlEnvVarName, rackhdUrl)
	os.Setenv(AnsibleRackHdConfigPath, configPath)
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
	// Unsetting environment variables
	os.Unsetenv(RackHdApiUrlEnvVarName)
	os.Unsetenv(AnsibleRackHdConfigPath)
}

//TODO
func TestHandleList(t *testing.T){

}

//TODO
func TestHandleHost(t *testing.T){

}