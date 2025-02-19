package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/rootlyhq/terraform-provider-rootly/client"
	"github.com/rootlyhq/terraform-provider-rootly/tools"
)

func resourceService() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceServiceCreate,
		ReadContext:   resourceServiceRead,
		UpdateContext: resourceServiceUpdate,
		DeleteContext: resourceServiceDelete,
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
				Description: "The name of the service",
			},

			"slug": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The slug of the service",
			},

			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The description of the service",
			},

			"public_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The public description of the service",
			},

			"notify_emails": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: tools.EqualIgnoringOrder,
				Computed:         true,
				Required:         false,
				Optional:         true,
				Description:      "Emails attached to the service",
			},

			"color": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The hex color of the service",
			},

			"status": &schema.Schema{
				Type:        schema.TypeString,
				Default:     "operational",
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The status of the service. Value must be one of `operational`, `impacted`, `outage`, `partial_outage`, `major_outage`.",
			},

			"position": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "Position of the service",
			},

			"backstage_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The Backstage entity id associated to this service. eg: :namespace/:kind/:entity_name",
			},

			"pagerduty_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The PagerDuty service id associated to this service",
			},

			"opsgenie_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The Opsgenie service id associated to this service",
			},

			"github_repository_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The GitHub repository name associated to this service. eg: rootlyhq/my-service",
			},

			"github_repository_branch": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The GitHub repository branch associated to this service. eg: main",
			},

			"gitlab_repository_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The Gitlab repository name associated to this service. eg: rootlyhq/my-service",
			},

			"gitlab_repository_branch": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Required:    false,
				Optional:    true,
				ForceNew:    false,
				Description: "The Gitlab repository branch associated to this service. eg: main",
			},

			"environment_ids": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: tools.EqualIgnoringOrder,
				Computed:         true,
				Required:         false,
				Optional:         true,
				Description:      "Environments associated with this service",
			},

			"service_ids": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: tools.EqualIgnoringOrder,
				Computed:         true,
				Required:         false,
				Optional:         true,
				Description:      "Services dependent on this service",
			},

			"owners_group_ids": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				DiffSuppressFunc: tools.EqualIgnoringOrder,
				Computed:         true,
				Required:         false,
				Optional:         true,
				Description:      "Owner Teams associated with this service",
			},

			"owners_user_ids": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "Owner Users associated with this service",
			},

			"slack_channels": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Required:    false,
				Optional:    true,
				Description: "Slack Channels associated with this service",
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
				Description: "Slack Aliases associated with this service",
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

func resourceServiceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)

	tflog.Trace(ctx, fmt.Sprintf("Creating Service"))

	s := &client.Service{}

	if value, ok := d.GetOkExists("name"); ok {
		s.Name = value.(string)
	}
	if value, ok := d.GetOkExists("slug"); ok {
		s.Slug = value.(string)
	}
	if value, ok := d.GetOkExists("description"); ok {
		s.Description = value.(string)
	}
	if value, ok := d.GetOkExists("public_description"); ok {
		s.PublicDescription = value.(string)
	}
	if value, ok := d.GetOkExists("notify_emails"); ok {
		s.NotifyEmails = value.([]interface{})
	}
	if value, ok := d.GetOkExists("color"); ok {
		s.Color = value.(string)
	}
	if value, ok := d.GetOkExists("status"); ok {
		s.Status = value.(string)
	}
	if value, ok := d.GetOkExists("position"); ok {
		s.Position = value.(int)
	}
	if value, ok := d.GetOkExists("backstage_id"); ok {
		s.BackstageId = value.(string)
	}
	if value, ok := d.GetOkExists("pagerduty_id"); ok {
		s.PagerdutyId = value.(string)
	}
	if value, ok := d.GetOkExists("opsgenie_id"); ok {
		s.OpsgenieId = value.(string)
	}
	if value, ok := d.GetOkExists("github_repository_name"); ok {
		s.GithubRepositoryName = value.(string)
	}
	if value, ok := d.GetOkExists("github_repository_branch"); ok {
		s.GithubRepositoryBranch = value.(string)
	}
	if value, ok := d.GetOkExists("gitlab_repository_name"); ok {
		s.GitlabRepositoryName = value.(string)
	}
	if value, ok := d.GetOkExists("gitlab_repository_branch"); ok {
		s.GitlabRepositoryBranch = value.(string)
	}
	if value, ok := d.GetOkExists("environment_ids"); ok {
		s.EnvironmentIds = value.([]interface{})
	}
	if value, ok := d.GetOkExists("service_ids"); ok {
		s.ServiceIds = value.([]interface{})
	}
	if value, ok := d.GetOkExists("owners_group_ids"); ok {
		s.OwnersGroupIds = value.([]interface{})
	}
	if value, ok := d.GetOkExists("owners_user_ids"); ok {
		s.OwnersUserIds = value.([]interface{})
	}
	if value, ok := d.GetOkExists("slack_channels"); ok {
		s.SlackChannels = value.([]interface{})
	}
	if value, ok := d.GetOkExists("slack_aliases"); ok {
		s.SlackAliases = value.([]interface{})
	}

	res, err := c.CreateService(s)
	if err != nil {
		return diag.Errorf("Error creating service: %s", err.Error())
	}

	d.SetId(res.ID)
	tflog.Trace(ctx, fmt.Sprintf("created a service resource: %s", d.Id()))

	return resourceServiceRead(ctx, d, meta)
}

func resourceServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Reading Service: %s", d.Id()))

	item, err := c.GetService(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("Service (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}

		return diag.Errorf("Error reading service: %s", d.Id())
	}

	d.Set("name", item.Name)
	d.Set("slug", item.Slug)
	d.Set("description", item.Description)
	d.Set("public_description", item.PublicDescription)
	d.Set("notify_emails", item.NotifyEmails)
	d.Set("color", item.Color)
	d.Set("status", item.Status)
	d.Set("position", item.Position)
	d.Set("backstage_id", item.BackstageId)
	d.Set("pagerduty_id", item.PagerdutyId)
	d.Set("opsgenie_id", item.OpsgenieId)
	d.Set("github_repository_name", item.GithubRepositoryName)
	d.Set("github_repository_branch", item.GithubRepositoryBranch)
	d.Set("gitlab_repository_name", item.GitlabRepositoryName)
	d.Set("gitlab_repository_branch", item.GitlabRepositoryBranch)
	d.Set("environment_ids", item.EnvironmentIds)
	d.Set("service_ids", item.ServiceIds)
	d.Set("owners_group_ids", item.OwnersGroupIds)
	d.Set("owners_user_ids", item.OwnersUserIds)
	d.Set("slack_channels", item.SlackChannels)
	d.Set("slack_aliases", item.SlackAliases)

	return nil
}

func resourceServiceUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Updating Service: %s", d.Id()))

	s := &client.Service{}

	if d.HasChange("name") {
		s.Name = d.Get("name").(string)
	}
	if d.HasChange("slug") {
		s.Slug = d.Get("slug").(string)
	}
	if d.HasChange("description") {
		s.Description = d.Get("description").(string)
	}
	if d.HasChange("public_description") {
		s.PublicDescription = d.Get("public_description").(string)
	}
	if d.HasChange("notify_emails") {
		s.NotifyEmails = d.Get("notify_emails").([]interface{})
	}
	if d.HasChange("color") {
		s.Color = d.Get("color").(string)
	}
	if d.HasChange("status") {
		s.Status = d.Get("status").(string)
	}
	if d.HasChange("position") {
		s.Position = d.Get("position").(int)
	}
	if d.HasChange("backstage_id") {
		s.BackstageId = d.Get("backstage_id").(string)
	}
	if d.HasChange("pagerduty_id") {
		s.PagerdutyId = d.Get("pagerduty_id").(string)
	}
	if d.HasChange("opsgenie_id") {
		s.OpsgenieId = d.Get("opsgenie_id").(string)
	}
	if d.HasChange("github_repository_name") {
		s.GithubRepositoryName = d.Get("github_repository_name").(string)
	}
	if d.HasChange("github_repository_branch") {
		s.GithubRepositoryBranch = d.Get("github_repository_branch").(string)
	}
	if d.HasChange("gitlab_repository_name") {
		s.GitlabRepositoryName = d.Get("gitlab_repository_name").(string)
	}
	if d.HasChange("gitlab_repository_branch") {
		s.GitlabRepositoryBranch = d.Get("gitlab_repository_branch").(string)
	}
	if d.HasChange("environment_ids") {
		s.EnvironmentIds = d.Get("environment_ids").([]interface{})
	}
	if d.HasChange("service_ids") {
		s.ServiceIds = d.Get("service_ids").([]interface{})
	}
	if d.HasChange("owners_group_ids") {
		s.OwnersGroupIds = d.Get("owners_group_ids").([]interface{})
	}
	if d.HasChange("owners_user_ids") {
		s.OwnersUserIds = d.Get("owners_user_ids").([]interface{})
	}
	if d.HasChange("slack_channels") {
		s.SlackChannels = d.Get("slack_channels").([]interface{})
	}
	if d.HasChange("slack_aliases") {
		s.SlackAliases = d.Get("slack_aliases").([]interface{})
	}

	_, err := c.UpdateService(d.Id(), s)
	if err != nil {
		return diag.Errorf("Error updating service: %s", err.Error())
	}

	return resourceServiceRead(ctx, d, meta)
}

func resourceServiceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*client.Client)
	tflog.Trace(ctx, fmt.Sprintf("Deleting Service: %s", d.Id()))

	err := c.DeleteService(d.Id())
	if err != nil {
		// In the case of a NotFoundError, it means the resource may have been removed upstream.
		// We just remove it from the state.
		if _, ok := err.(client.NotFoundError); ok && !d.IsNewResource() {
			tflog.Warn(ctx, fmt.Sprintf("Service (%s) not found, removing from state", d.Id()))
			d.SetId("")
			return nil
		}
		return diag.Errorf("Error deleting service: %s", err.Error())
	}

	d.SetId("")

	return nil
}
