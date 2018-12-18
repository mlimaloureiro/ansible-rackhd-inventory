package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/mlimaloureiro/ansible-rackhd-inventory/rackhd"
	"github.com/spf13/viper"
	"os"
)

const (
	RackHdApiUrlEnvVarName  = "RACK_HD_API_URL"
	AnsibleRackHdConfigPath = "ANSIBLE_RACKHD_CONFIG_PATH"
)

type props struct {
	rackhdUrl   string
	groups      []string
	filterGroup string
}

type CLIArgs struct {
	List bool
	Host string
}

func main() {
	props := getPropsFromConfig()
	args := CLIArgs{}
	arg.MustParse(&args)
	if args.List {
		output, err := handleList(props)
		if err != nil {
			panic(fmt.Errorf("Fatal error handling list: %s \n", err))
		}
		marshalResult, _ := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(marshalResult))

		return
	}

}

func getPropsFromConfig() props {
	config := viper.New()

	envRackhdApiUrl, envRackhdApiUrlOk := os.LookupEnv(RackHdApiUrlEnvVarName)
	envAnsibleRackhdConfigPath, envAnsibleRackhdConfigPathOk := os.LookupEnv(AnsibleRackHdConfigPath)
	if envAnsibleRackhdConfigPathOk {
		config.SetConfigFile(envAnsibleRackhdConfigPath)
	} else {
		config.AddConfigPath(".")
		config.SetConfigName("config")
		config.SetConfigType("yml")
	}
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	var rackhdUrl string
	if envRackhdApiUrlOk {
		rackhdUrl = envRackhdApiUrl
	} else {
		rackhdUrl = config.GetString("rackhd_api_url")
	}

	return props{
		rackhdUrl: rackhdUrl, groups: config.GetStringSlice("groups"),
		filterGroup: config.GetString("filter_group")}
}

func getGroupNodesAndVars(props props) (map[string]interface{}, Hostvars, error) {
	rackhdClient := rackhd.Client{BaseUrl: props.rackhdUrl}

	hostvars := Hostvars{}
	groups := make(map[string]interface{})
	if len(props.groups) == 0 {
		var err error
		props.groups, err = rackhdClient.GetAllTags()
		if err != nil {
			return nil, nil, err
		}
	}

	for _, group := range props.groups {
		result, err := rackhdClient.GetTaggedNodesIpAddress(group)
		if err != nil {
			return groups, nil, err
		}
		hostsFoundByTag := len(result) == 0
		if !hostsFoundByTag {
			groupItem := GroupItem{}
			for _, item := range result {
				hostvars[item] = HostvarsItem{
					AnsibleSSHHostPrivate: item,
					AnsibleSSHHost:        item,
				}
				groupItem.Hosts = append(groupItem.Hosts, item)
			}
			groups[group] = groupItem
		}
	}

	return groups, hostvars, nil
}

func handleList(props props) (map[string]interface{}, error) {
	output := make(map[string]interface{})
	groups, hostvars, err := getGroupNodesAndVars(props)
	if err != nil {

		return nil, err
	}

	if props.filterGroup != "" {
		props.groups = []string{props.filterGroup}
		filterGroup, _, err := getGroupNodesAndVars(props)

		if err != nil {
			return nil, err
		}

		filterGroupItem, _ := filterGroup[props.filterGroup].(GroupItem)
		for groupName, groupItem := range groups {
			groupItem, _ := groupItem.(GroupItem)
			intersectionOfGroups := IntersectionOfTwoSlices(groupItem.Hosts, filterGroupItem.Hosts)
			if len(intersectionOfGroups) > 0 {
				groups[groupName] =
					GroupItem{
						Hosts: intersectionOfGroups}
			}
		}

		for hostname := range hostvars {
			if !ValueInSlice(filterGroupItem.Hosts, hostname) {
				delete(hostvars, hostname)
			}
		}

	}

	for groupName, groupItem := range groups {
		output[groupName] = groupItem
	}
	output["_meta"] = Meta{
		Hostvars: hostvars,
	}

	return output, err
}
