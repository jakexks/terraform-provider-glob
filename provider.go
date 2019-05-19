package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider contains the definitions that make this a provider
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"glob_list": resourceGlob(),
		},
	}
}
