package main

import (
	"github.com/mlimaloureiro/ansible-rackhd-inventory/rackhd"
	"encoding/json"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/spf13/viper"
)

type props struct {
	rackhdUrl string
	groups    []string
}

type CLIArgs struct {
	List bool
}

func main() {
	config := viper.New()
	config.AddConfigPath(".")
	config.SetConfigName("config")
	config.SetConfigType("yml")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	props := props{rackhdUrl: config.GetString("rackhd_api_url"), groups: config.GetStringSlice("groups")}

	args := CLIArgs{}
	arg.MustParse(&args)
	if args.List {
		handleList(props)
		return
	}

	handleList(props)

	return
}

func handleList(props props) {
	rackhdClient := rackhd.Client{BaseUrl: props.rackhdUrl}

	output := make(map[string]interface{})
	hostvars := Hostvars{}

	for _, group := range props.groups {
		result, err := rackhdClient.GetTaggedNodesIpAddress(group)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(result) == 0 {
			continue
		}

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

	output["_meta"] = Meta{
		Hostvars: hostvars,
	}
	marshalResult, _ := json.MarshalIndent(output, "", "  ")
	fmt.Println(string(marshalResult))
}
