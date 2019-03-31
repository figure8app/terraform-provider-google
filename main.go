package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/odedniv/terraform-provider-google/google"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: google.Provider})
}
