package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider contains the definitions that make this a provider
func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"glob_filename_list": dataSourceGlobFilenameList(),
			"glob_contents_list": dataSourceGlobContentsList(),
			"glob_map":           dataSourceGlobMap(),
		},
	}
}
