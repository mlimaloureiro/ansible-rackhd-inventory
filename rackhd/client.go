package rackhd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	tagPathTemplate = "/api/current/tags/%s/nodes"
	lookupPath      = "/api/2.0/lookups"
	tagsPath        = "/api/current/tags"
)

// Client config
type Client struct {
	// BaseUrl has the base url for rackhd
	BaseUrl string
}

// ResponseCollection holds RHD response
type NodesWithTag []ResponseIdItem

// ResponseItem
type ResponseIdItem struct {
	ID   string `json:"id"`
	// name in response equals to the machine mac address
	MacAddress string `json:"name"`
}

// LookupTable holds a list of node ids and associated ip addresses
type LookupTable []LookupItem

// LookupItem has the node id and ip address
type LookupItem struct {
	ID         string `json:"id"`
	Node       string `json:"node"`
	IpAddress  string `json:"ipAddress"`
	MacAddress string `json:"macAddress"`
}

// TagList holds a list of all tags
type TagsList []string

// GetTaggedNodesIpAddress find nodes with specific tag
func (c Client) GetTaggedNodesIpAddress(tagName string) ([]string, error) {
	nodes, err := c.fetchTaggedNodes(tagName)
	if err != nil {
		return []string{}, fmt.Errorf("error fetching nodes: %s", err)
	}

	lookupTable, err := c.fetchLookupTable()
	if err != nil {
		return []string{}, fmt.Errorf("error fetching lookup table: %s", err)
	}

	return c.filterIpAddresses(lookupTable, nodes), nil
}

func (c Client) GetAllTags() (TagsList, error) {
	tagRequest, err := c.request(tagsPath)
	if err != nil {
		return TagsList{}, err
	}
	var tagsList TagsList
	err = json.Unmarshal(tagRequest, &tagsList)
	if err != nil {
		return TagsList{}, err
	}

	return tagsList, nil
}

func (c Client) fetchTaggedNodes(tagName string) (NodesWithTag, error) {
	tagPath := fmt.Sprintf(tagPathTemplate, tagName)
	tagRequest, err := c.request(tagPath)
	if err != nil {
		return NodesWithTag{}, err
	}
	var nodes NodesWithTag
	err = json.Unmarshal(tagRequest, &nodes)
	if err != nil {
		return NodesWithTag{}, err
	}

	return nodes, nil
}

func (c Client) fetchLookupTable() (LookupTable, error) {
	lookupRequest, err := c.request(lookupPath)
	if err != nil {
		return LookupTable{}, err
	}
	var lookupTable LookupTable
	err = json.Unmarshal(lookupRequest, &lookupTable)
	if err != nil {
		return LookupTable{}, err
	}

	return lookupTable, nil
}

func (c Client) request(path string) ([]byte, error) {
	response, err := http.Get(fmt.Sprintf("%s%s", c.BaseUrl, path))
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, err
}

func (c Client) filterIpAddresses(lookupTable LookupTable, nodes NodesWithTag) []string {
	var ipAddresses []string
	var isActiveNode bool
	for _, node := range nodes {
		for _, lookupItem := range lookupTable {
			isActiveNode = lookupItem.IpAddress != "" &&
				lookupItem.Node == node.ID &&
				lookupItem.MacAddress == node.MacAddress

			if isActiveNode {
				ipAddresses = append(ipAddresses, lookupItem.IpAddress)
			}
		}
	}

	return ipAddresses
}
