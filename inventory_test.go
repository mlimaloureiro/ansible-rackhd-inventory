package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

const (
	rackhdUrl  = "http://127.0.0.1:5000"
	configPath = "./config.test.yaml"
)

func TestConfigReads(t *testing.T) {
	var props props
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

func TestHandleList(t *testing.T) {
	err := os.Setenv(AnsibleRackHdConfigPath, configPath)
	if err != nil {
		t.Errorf("%s\n", err)

		return
	}
	err = os.Setenv(AnsibleRackHdConfigPath, configPath)
	if err != nil {
		t.Errorf("%s\n", err)

		return
	}
	props := getPropsFromConfig()
	output, err := handleList(props)
	if value, ok := output["all"]; !ok {
		t.Errorf("Expected key 'all' got %s\n", value)
	}
	if value, ok := output["ungrouped"]; !ok {
		t.Errorf("Expected key 'ungrouped' got %s\n", value)
	}
	marshalResult, _ := json.MarshalIndent(output, "", "  ")
	fmt.Println(string(marshalResult))
	err = os.Unsetenv(RackHdApiUrlEnvVarName)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	err = os.Unsetenv(AnsibleRackHdConfigPath)
	if err != nil {
		t.Errorf("%s\n", err)
	}
}
