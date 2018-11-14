package main

// Meta holds every hostvars info
type Meta struct {
	Hostvars Hostvars `json:"hostvars"`
}

// Hostvars hold a key value with host information
type Hostvars map[string]HostvarsItem

// HostvarsItem holds the values for each host
type HostvarsItem struct {
	AnsibleSSHHost        string `json:"ansible_ssh_host"`
	AnsibleSSHHostPrivate string `json:"ansible_ssh_host_private"`
}

// GroupItem has the information of each specific group
type GroupItem struct {
	Hosts []string `json:"hosts"`
}
