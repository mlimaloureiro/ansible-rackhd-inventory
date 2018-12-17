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

// getPropsFromConfig returns a props instance with values from config.yml
func getPropsFromConfig() props {
	config := viper.New()

	envRackhdApiUrl, envRackhdApiUrlOk := os.LookupEnv(RackHdApiUrlEnvVarName)
	envAnsibleRackhdConfigPath, envAnsibleRackhdConfigPathOk := os.LookupEnv(AnsibleRackHdConfigPath)
	if envAnsibleRackhdConfigPathOk == true {
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
	if envRackhdApiUrlOk == true {
		rackhdUrl = envRackhdApiUrl
	} else {
		rackhdUrl = config.GetString("rackhd_api_url")
	}

	return props{
		rackhdUrl: rackhdUrl, groups: config.GetStringSlice("groups"),
		filterGroup: config.GetString("filter_group")}
}

func handleList(props props) (map[string]interface{}, error) {
	rackhdClient := rackhd.Client{BaseUrl: props.rackhdUrl}

	hostvars := Hostvars{}
	output := make(map[string]interface{})
	for _, group := range props.groups {
		result, err := rackhdClient.GetTaggedNodesIpAddress(group)
		if err != nil {
			return output, err
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
			output[group] = groupItem
		}
	}

	output["_meta"] = Meta{
		Hostvars: hostvars,
	}

	return output, nil
}
