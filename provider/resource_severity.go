package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rootlyhq/terraform-provider-rootly/client"
)

func resourceSeverity() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSeverityCreate,
		ReadContext:   resourceSeverityRead,
		UpdateContext: resourceSeverityUpdate,
		DeleteContext: resourceSeverityDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{

			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    false,
				Required:    true,
				Optional:    false,
				ForceNew:    false,
				Description: "The name of the severity",
			},

			"slug": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The slug of the severity",
			},

			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The description of the severity",
			},

			"severity": &schema.Schema{
				Type:        schema.TypeString,
				Default:     "critical",
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The severity of the severity. Value must be one of `critical`, `high`, `medium`, `low`.",
			},

			"color": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "",
			},

			"notify_emails": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "Emails to attach to the severity",
			},

			"slack_channels": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "Slack Channels associated with this severity",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},

			"slack_aliases": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "Slack Aliases associated with this severity",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceSeverityCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	tflog.Trace(ctx, fmt.Sprintf("Creating Severity"))

	s := &client.Severity{}

	if value, ok := d.GetOkExists("name"); ok {
		s.Name = value.(string)
	}
	if value, ok := d.GetOkExists("slug"); ok {
		s.Slug = value.(string)
	}
	if value, ok := d.GetOkExists("description"); ok {
		s.Description = value.(string)
	}
	if value, ok := d.GetOkExists("severity"); ok {
		s.Severity = value.(string)
	}
	if value, ok := d.GetOkExists("color"); ok {
		s.Color = value.(string)
	}
	if value, ok := d.GetOkExists("notify_emails"); ok {
		s.NotifyEmails = value.([]interface{})
	}
	if value, ok := d.GetOkExists("slack_channels"); ok {
		s.SlackChannels = value.([]interface{})
	}
	if value, ok := d.GetOkExists("slack_aliases"); ok {
		s.SlackAliases = value.([]interface{})
	}

	res, err := c.CreateSeverity(s)
	if err != nil {
		return diag.Errorf("Error creating severity: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created a severity resource: %s", d.Id()))

	return resourceSeverityRead(ctx, d, meta)
}

func resourceSeverityRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading Severity: %s", d.Id()))

	item, err := c.GetSeverity(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("Severity (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading severity: %s", d.Id())
	}

	d.Set("name", item.Name)
	d.Set("slug", item.Slug)
	d.Set("description", item.Description)
	d.Set("severity", item.Severity)
	d.Set("color", item.Color)
	d.Set("notify_emails", item.NotifyEmails)
	d.Set("slack_channels", item.SlackChannels)
	d.Set("slack_aliases", item.SlackAliases)

	return nil
}

func resourceSeverityUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating Severity: %s", d.Id()))

	s := &client.Severity{}

	if d.HasChange("name") {
		s.Name = d.Get("name").(string)
	}
	if d.HasChange("slug") {
		s.Slug = d.Get("slug").(string)
	}
	if d.HasChange("description") {
		s.Description = d.Get("description").(string)
	}
	if d.HasChange("severity") {
		s.Severity = d.Get("severity").(string)
	}
	if d.HasChange("color") {
		s.Color = d.Get("color").(string)
	}
	if d.HasChange("notify_emails") {
		s.NotifyEmails = d.Get("notify_emails").([]interface{})
	}
	if d.HasChange("slack_channels") {
		s.SlackChannels = d.Get("slack_channels").([]interface{})
	}
	if d.HasChange("slack_aliases") {
		s.SlackAliases = d.Get("slack_aliases").([]interface{})
	}

	_, err := c.UpdateSeverity(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating severity: %s", err.Error())
	}

	return resourceSeverityRead(ctx, d, meta)
}

func resourceSeverityDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting Severity: %s", d.Id()))

	err := c.DeleteSeverity(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("Severity (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting severity: %s", err.Error())
	}

	d.SetId("")

	return nil
}
