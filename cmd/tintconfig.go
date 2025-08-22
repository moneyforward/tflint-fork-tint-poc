package cmd

import (
	"os"
	"path/filepath"
)

var autoDiscoveryPaths = []string{
	"services-terraform",
	"src/services-terraform",
	"projects/services-terraform",
	"work/services-terraform",
	"Documents/services-terraform",
	"dev/services-terraform",
	"moneyforward/services-terraform",
}

func InjectTintConfig(opts *Options) {
	if opts.ActAsBundledPlugin || opts.ActAsWorker {
		return
	}

	terraformRoot := opts.TerraformRoot
	if terraformRoot == "" {
		terraformRoot = discoverServicesTerraformRoot()
	}
	serviceDir := opts.Service
	if opts.Chdir == "" && serviceDir == "" {
		panic("Service directory or chdir is required") // TODO: Fix error handling
	}
	opts.Chdir = filepath.Join(terraformRoot, "workload", serviceDir) // TODO: Decide whether to include 'workload'
}

func discoverServicesTerraformRoot() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return ""
	}

	for _, relativePath := range autoDiscoveryPaths {
		path := filepath.Join(homeDir, relativePath)
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return ""
}
