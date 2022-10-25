package provider

import (
	"context"
	
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rootlyhq/terraform-provider-rootly/client"
	rootlygo "github.com/rootlyhq/terraform-provider-rootly/schema"
)

func dataSourceEnvironment() *schema.Resource{
	return &schema.Resource{
		ReadContext: dataSourceEnvironmentRead,
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
			

			"color": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
				Optional: true,
			},
			

				"created_at": &schema.Schema{
					Type: schema.TypeMap,
					Description: "Filter by date range using 'lt' and 'gt'.",
					Optional: true,
				},
				
		},
	}
}

func dataSourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	params := new(rootlygo.ListEnvironmentsParams)
	page_size := 1
	params.PageSize = &page_size

	
				if value, ok := d.GetOkExists("slug"); ok {
					slug := value.(string)
					params.FilterSlug = &slug
				}
			

				if value, ok := d.GetOkExists("name"); ok {
					name := value.(string)
					params.FilterName = &name
				}
			

				if value, ok := d.GetOkExists("color"); ok {
					color := value.(string)
					params.FilterColor = &color
				}
			

				created_at_gt := d.Get("created_at").(map[string]interface{})
				if value, exists := created_at_gt["gt"]; exists {
					v := value.(string)
					params.FilterCreatedAtGt = &v
				}
			

				created_at_lt := d.Get("created_at").(map[string]interface{})
				if value, exists := created_at_lt["lt"]; exists {
					v := value.(string)
					params.FilterCreatedAtLt = &v
				}
			

	items, err := c.ListEnvironments(params)
	if err != nil {
		return diag.FromErr(err)
	}

	if len(items) == 0 {
		return diag.Errorf("environment not found")
	}
	item, _ := items[0].(*client.Environment)

	d.SetId(item.ID)

	return nil
}
