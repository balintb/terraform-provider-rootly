package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rootlyhq/terraform-provider-rootly/client"
)

func resourcePlaybook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePlaybookCreate,
		ReadContext:   resourcePlaybookRead,
		UpdateContext: resourcePlaybookUpdate,
		DeleteContext: resourcePlaybookDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			"title": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    false,
				Required:    true,
				Optional:    false,
				ForceNew:    false,
				Description: "The title of the playbook",
			},

			"summary": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The summary of the playbook",
			},

			"external_url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The external url of the playbook",
			},

			"severity_ids": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "The Severity ID's to attach to the incident",
			},

			"environment_ids": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "The Environment ID's to attach to the incident",
			},

			"functionality_ids": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "The Functionality ID's to attach to the incident",
			},

			"group_ids": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "The Team ID's to attach to the incident",
			},

			"incident_type_ids": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "The Incident Type ID's to attach to the incident",
			},
		},
	}
}

func resourcePlaybookCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	tflog.Trace(ctx, fmt.Sprintf("Creating Playbook"))

	s := &client.Playbook{}

	if value, ok := d.GetOkExists("title"); ok {
		s.Title = value.(string)
	}
	if value, ok := d.GetOkExists("summary"); ok {
		s.Summary = value.(string)
	}
	if value, ok := d.GetOkExists("external_url"); ok {
		s.ExternalUrl = value.(string)
	}
	if value, ok := d.GetOkExists("severity_ids"); ok {
		s.SeverityIds = value.([]interface{})
	}
	if value, ok := d.GetOkExists("environment_ids"); ok {
		s.EnvironmentIds = value.([]interface{})
	}
	if value, ok := d.GetOkExists("functionality_ids"); ok {
		s.FunctionalityIds = value.([]interface{})
	}
	if value, ok := d.GetOkExists("group_ids"); ok {
		s.GroupIds = value.([]interface{})
	}
	if value, ok := d.GetOkExists("incident_type_ids"); ok {
		s.IncidentTypeIds = value.([]interface{})
	}

	res, err := c.CreatePlaybook(s)
	if err != nil {
		return diag.Errorf("Error creating playbook: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created a playbook resource: %s", d.Id()))

	return resourcePlaybookRead(ctx, d, meta)
}

func resourcePlaybookRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading Playbook: %s", d.Id()))

	item, err := c.GetPlaybook(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("Playbook (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading playbook: %s", d.Id())
	}

	d.Set("title", item.Title)
	d.Set("summary", item.Summary)
	d.Set("external_url", item.ExternalUrl)
	d.Set("severity_ids", item.SeverityIds)
	d.Set("environment_ids", item.EnvironmentIds)
	d.Set("functionality_ids", item.FunctionalityIds)
	d.Set("group_ids", item.GroupIds)
	d.Set("incident_type_ids", item.IncidentTypeIds)

	return nil
}

func resourcePlaybookUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating Playbook: %s", d.Id()))

	s := &client.Playbook{}

	if d.HasChange("title") {
		s.Title = d.Get("title").(string)
	}
	if d.HasChange("summary") {
		s.Summary = d.Get("summary").(string)
	}
	if d.HasChange("external_url") {
		s.ExternalUrl = d.Get("external_url").(string)
	}
	if d.HasChange("severity_ids") {
		s.SeverityIds = d.Get("severity_ids").([]interface{})
	}
	if d.HasChange("environment_ids") {
		s.EnvironmentIds = d.Get("environment_ids").([]interface{})
	}
	if d.HasChange("functionality_ids") {
		s.FunctionalityIds = d.Get("functionality_ids").([]interface{})
	}
	if d.HasChange("group_ids") {
		s.GroupIds = d.Get("group_ids").([]interface{})
	}
	if d.HasChange("incident_type_ids") {
		s.IncidentTypeIds = d.Get("incident_type_ids").([]interface{})
	}

	_, err := c.UpdatePlaybook(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating playbook: %s", err.Error())
	}

	return resourcePlaybookRead(ctx, d, meta)
}

func resourcePlaybookDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting Playbook: %s", d.Id()))

	err := c.DeletePlaybook(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("Playbook (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting playbook: %s", err.Error())
	}

	d.SetId("")

	return nil
}
