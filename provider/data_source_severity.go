package provider

import (
	"context"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rootlyhq/terraform-provider-rootly/client"
	rootlygo "github.com/rootlyhq/terraform-provider-rootly/schema"
)

func dataSourceSeverity() *schema.Resource{
	return &schema.Resource{
		ReadContext: dataSourceSeverityRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
			},
			
			"name": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
				Optional: true,
			},
			

			"slug": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
				Optional: true,
			},
			

			"severity": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
				Optional: true,
			},
			

			"color": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
				Optional: true,
			},
			
		},
	}
}

func dataSourceSeverityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	params := new(rootlygo.ListSeveritiesParams)
	page_size := 1
	params.PageSize = &page_size

	
				slug := d.Get("slug").(string)
				params.FilterSlug = &slug
			

				name := d.Get("name").(string)
				params.FilterName = &name
			

				severity := d.Get("severity").(string)
				params.FilterSeverity = &severity
			

				color := d.Get("color").(string)
				params.FilterColor = &color
			

	items, err := c.ListSeverities(params)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(items) == 0 {
		return diag.Errorf("severity not found")
	}
	item, _ := items[0].(*client.Severity)

	d.SetId(item.ID)

	return nil
}
