package main

import (
	"net/http/httptest"
	"os"
	"testing"
)

const (
	configPath = "./config.test.yaml"
)

func setEnvironmentVars(rackhdUrl string, configPath string) error {
	err := os.Setenv(RackHdApiUrlEnvVarName, rackhdUrl)
	if err != nil {

		return err
	}
	err = os.Setenv(AnsibleRackHdConfigPath, configPath)
	if err != nil {

		return err
	}

	return nil
}

func TestConfigReads(t *testing.T) {
	server := httptest.NewServer(RackhdPathHandlers())
	defer server.Close()
	err := setEnvironmentVars(server.URL, configPath)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	props := getPropsFromConfig()
	if props.rackhdUrl != server.URL {
		t.Errorf("\n%s  \n%s", props.rackhdUrl, server.URL)
	}
	if len(props.groups) != 3 {
		t.Errorf("\n%d  \n%d", len(props.groups), 3)
	}
	if props.groups[1] != TagCephNode {
		t.Errorf("\n%s  \n%s", props.groups[2], TagCephNode)
	}
	if props.filterGroup != TagNew {
		t.Errorf("\n%s  \n%s", props.filterGroup, TagNew)
	}

}

func TestHandleList(t *testing.T) {
	server := httptest.NewServer(RackhdPathHandlers())
	defer server.Close()
	err := setEnvironmentVars(server.URL, configPath)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	props := getPropsFromConfig()
	output, err := handleList(props)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	if value, ok := output["all"]; !ok {
		t.Errorf("Expected key 'all' got %s\n", value)
	}
	if value, ok := output["ungrouped"]; !ok {
		t.Errorf("Expected key 'ungrouped' got %s\n", value)
	}

}

func TestHandleHost(t *testing.T) {
	server := httptest.NewServer(RackhdPathHandlers())
	defer server.Close()
	const hostname = "192.168.1.130"
	err := setEnvironmentVars(server.URL, configPath)
	if err != nil {
		t.Errorf("%s\n", err)
	}
	props := getPropsFromConfig()

	hostvarItem, err := handleHost(hostname, props)
	if hostvarItem.AnsibleSSHHostPrivate != hostname {
		t.Errorf("\n%s \n%s", hostvarItem.AnsibleSSHHostPrivate, hostname)
	}
	if hostvarItem.AnsibleSSHHost != hostname {
		t.Errorf("\n%s \n%s", hostvarItem.AnsibleSSHHost, hostname)
	}

}
