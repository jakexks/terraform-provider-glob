package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceGlobMap() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGlobMapRead,
		Schema: map[string]*schema.Schema{
			"pattern": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"matches": &schema.Schema{
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func dataSourceGlobMapRead(d *schema.ResourceData, m interface{}) error {
	p := d.Get("pattern").(string)
	matches, err := filepath.Glob(p)
	if err != nil {
		return err
	}
	sort.Strings(matches)

	var contents []string
	var allcontents string
	for _, file := range matches {
		allcontents += file
		f, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}
		contents = append(contents, string(f))
		allcontents += string(f)
	}
	filemap := make(map[string]interface{})
	for i, match := range matches {
		filemap[match] = contents[i]
	}

	d.SetId(fmt.Sprintf("%x", sha256.Sum256([]byte(allcontents))))
	d.Set("matches", filemap)
	return nil
}
