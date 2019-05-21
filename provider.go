package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider contains the definitions that make this a provider
func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"glob_list": dataSourceGlobList(),
			"glob_map":  dataSourceGlobMap(),
		},
	}
}
