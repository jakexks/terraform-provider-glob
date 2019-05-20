package main

import (
	"crypto/sha256"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGlob() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGlobRead,
		Schema: map[string]*schema.Schema{
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"matches": &schema.Schema{
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func dataSourceGlobRead(d *schema.ResourceData, m interface{}) error {
	p := d.Get("pattern").(string)
	items, err := filepath.Glob(p)
	if err != nil {
		return err
	}
	sort.Strings(items)
	var elemhash string
	for i, s := range items {
		elemhash += strconv.Itoa(i) + s
	}
	d.SetId(fmt.Sprintf("%x", sha256.Sum256([]byte(elemhash))))
	d.Set("matches", items)
	return nil
}
