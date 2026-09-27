package config

type Config struct {
	Sessions []Session `hcl:"session,block"`
}

type Session struct {
	Name    string   `hcl:"name,label"`
	Windows []Window `hcl:"window,block"`
	Focus   string   `hcl:"focus,optional"`
}

type Window struct {
	Name string   `hcl:"name,label"`
	Run  []string `hcl:"run,optional"`
	Send []string `hcl:"send,optional"`
}
