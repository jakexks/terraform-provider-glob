package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGlobContentsList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGlobContentsListRead,
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

func dataSourceGlobContentsListRead(d *schema.ResourceData, m interface{}) error {
	p := d.Get("pattern").(string)
	items, err := filepath.Glob(p)
	if err != nil {
		return err
	}
	sort.Strings(items)
	var totalcontents string
	var contents []string
	for _, s := range items {
		file, err := ioutil.ReadFile(s)
		if err != nil {
			return err
		}
		contents = append(contents, string(file))
		totalcontents += string(file)
	}
	d.SetId(fmt.Sprintf("%x", sha256.Sum256([]byte(totalcontents))))
	d.Set("matches", contents)
	return nil
}
