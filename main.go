package main

import (
	"github.com/CiscoInterCloudFabric/icf-terraform/icf"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: icf.Provider,
	})
}
