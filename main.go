package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/figure8app/terraform-provider-google/google"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: google.Provider})
}
