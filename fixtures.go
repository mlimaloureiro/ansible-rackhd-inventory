package main

const (
	RackHdApiUrlEnvVarName  = "RACK_HD_API_URL"
	AnsibleRackHdConfigPath = "ANSIBLE_RACKHD_CONFIG_PATH"
	EmptyString             = ""
	TwoSpaceString          = "  "
	EmptyBrackets           = "{}"
	EmptyBoxBrackets        = `[]`
	TagCephNode             = "ceph-node"
	TagCephMon              = "ceph-mon"
	TagNew                  = "new"
	lookupPath              = "/api/2.0/lookups"
	tagsPath                = "/api/current/tags"
	tagsPathTemplate        = "/api/current/tags/{tag}/nodes"
)
