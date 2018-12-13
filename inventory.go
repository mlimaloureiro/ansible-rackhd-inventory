package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexflint/go-arg"
	"github.com/mlimaloureiro/ansible-rackhd-inventory/rackhd"
	"github.com/spf13/viper"
	"os"
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
	// If argument "--list" was set to true
	if args.List {
		output, err := handleList(props)
		if err != nil {
			panic(fmt.Errorf("Fatal error handling list: %s \n", err))
		}
		marshalResult, _ := json.MarshalIndent(output, "", "  ")
		fmt.Println(string(marshalResult))
		// If argument "--host" was set instead of host.
	} else if args.Host != "" {
		output, err := handleHost(args.Host, props)
		if err != nil {
			panic(fmt.Errorf("Fatal error handling host: %s \n", err))
		}
		fmt.Printf("%+v\n", output)
	}
}

// getPropsFromConfig returns a props instance with values from config.yml
func getPropsFromConfig() props {
	// config allows us to set the path for the configuration yaml file and retrieve its values
	config := viper.New()

	// envRackhdApiUrl is the RackHD REST API URL defined by an environment variable and
	envRackhdApiUrl, envRackhdApiUrlOk := os.LookupEnv("RACK_HD_API_URL")
	envAnsibleRackhdConfigPath, envAnsibleRackhdConfigPathOk := os.LookupEnv("ANSIBLE_RACKHD_CONFIG_PATH")
	// If the "ANSIBLE_RACKHD_CONFIG_PATH" variable is set we define that as our config file
	if envAnsibleRackhdConfigPathOk == true {
		config.SetConfigFile(envAnsibleRackhdConfigPath)
		// Else we define the config.yaml as the default config file
	} else {
		config.AddConfigPath(".")
		config.SetConfigName("config")
		config.SetConfigType("yml")
	}
	// ReadConfig reads our config file
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// If the "RACK_HD_API_URL" is set we define that as our URL API
	var rackhdUrl string
	if envRackhdApiUrlOk == true {
		rackhdUrl = envRackhdApiUrl
	} else {
		rackhdUrl = config.GetString("rackhd_api_url")
	}

	return props{rackhdUrl:rackhdUrl, groups: config.GetStringSlice("groups"), filterGroup: config.GetString("filter_group")}
}

// handleList returns a map with tags as keys and the list of hosts as values
func handleList(props props) (map[string]interface{}, error) {
	// rackhdClient allows us to make calls to the REST API located at BaseUrl
	rackhdClient := rackhd.Client{BaseUrl: props.rackhdUrl}

	hostvars := Hostvars{}
	// ouput is used to return a map containing the host groups
	output := make(map[string]interface{})
	for _, group := range props.groups {
		// Makes a request to the RackHD api
		result, err := rackhdClient.GetTaggedNodesIpAddress(group)
		if err != nil {
			return output, err
		}
		// if there are no host for a given tag we jump to the next iteration
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
	return output, nil
}

// TODO: handleHost - Return hostvars values from this function
func handleHost(host string, props props) (map[string]interface{}, error) {
	fmt.Printf("%+v\n", props)
	// rackhdClient allows us to make calls to the REST API located at BaseUrl
	rackhdClient := rackhd.Client{BaseUrl: props.rackhdUrl}
	lookupTable, _ := rackhdClient.GetTaggedNodesIpAddress("")
	fmt.Printf("%+v\n", lookupTable)
	return nil, nil
}
