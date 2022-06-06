package renderer

import "github.com/hashicorp/terraform-config-inspect/tfconfig"

type Directory struct {
	Path    string
	Recurse bool
}

type Config struct {
	Directories []Directory
	OutputPath  string
}

type moduleData struct {
	Module   string
	Root     string
	TfModule *tfconfig.Module
}

type output struct {
	Config  Config
	Modules []moduleData
}
