package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGlob() *schema.Resource {
	return &schema.Resource{
		Create: resourceGlobCreate,
		Read:   resourceGlobRead,
		Update: resourceGlobUpdate,
		Delete: resourceGlobDelete,

		Schema: map[string]*schema.Schema{
			"directory": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceGlobCreate(d *schema.ResourceData, m interface{}) error {
	return resourceGlobRead(d, m)
}

func resourceGlobRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceGlobUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceGlobRead(d, m)
}

func resourceGlobDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
