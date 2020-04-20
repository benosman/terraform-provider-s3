package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/negronjl/terraform-provider-s3/s3"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: s3.Provider,
	})
}
