package main

import (
	"encoding/json"
	"testing"
)

func TestInterfacePayloadMarshalsCorrectly(t *testing.T) {
	output := make(map[string]interface{})
	output["_meta"] = Meta{
		Hostvars: Hostvars{
			"web1.dev": HostvarsItem{
				AnsibleSSHHost:        "127.0.0.1",
				AnsibleSSHHostPrivate: "127.0.0.1",
			},
			"web2.dev": HostvarsItem{
				AnsibleSSHHost:        "127.0.0.2",
				AnsibleSSHHostPrivate: "127.0.0.2",
			},
		},
	}
	output["web"] = GroupItem{
		Hosts: []string{
			"web1.dev",
			"web2.dev",
		},
	}

	marshalResult, _ := json.MarshalIndent(output, "", "	")
	if string(marshalResult) != expectedResult {
		t.Errorf("\n%s  \n%s", marshalResult, expectedResult)
	}
}

const expectedResult = `{
	"_meta": {
		"hostvars": {
			"web1.dev": {
				"ansible_ssh_host": "127.0.0.1",
				"ansible_ssh_host_private": "127.0.0.1"
			},
			"web2.dev": {
				"ansible_ssh_host": "127.0.0.2",
				"ansible_ssh_host_private": "127.0.0.2"
			}
		}
	},
	"web": {
		"hosts": [
			"web1.dev",
			"web2.dev"
		]
	}
}`
